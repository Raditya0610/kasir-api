package repository

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
)

type TransactionRepository interface {
	CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error)
	GetSalesReport(startDate, endDate string) (models.SalesReport, error)
}

type transactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *transactionRepository {
	return &transactionRepository{db: db}
}

func (repo *transactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		if stock < item.Quantity {
			return nil, fmt.Errorf("insufficient stock for product %s", productName)
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func (repo *transactionRepository) GetSalesReport(startDate, endDate string) (models.SalesReport, error) {
	var report models.SalesReport

	// 1. Total Revenue & Total Transaksi
	queryStats := `
		SELECT COALESCE(SUM(total_amount), 0), COUNT(id)
		FROM transactions
		WHERE created_at >= $1 AND created_at <= $2
	`
	err := repo.db.QueryRow(queryStats, startDate, endDate).Scan(&report.TotalRevenue, &report.TotalTransaksi)
	if err != nil {
		return report, err
	}

	// 2. Produk Terlaris
	queryBestSeller := `
		SELECT p.name, SUM(td.quantity) as qty
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE t.created_at >= $1 AND t.created_at <= $2
		GROUP BY p.name
		ORDER BY qty DESC
		LIMIT 1
	`
	err = repo.db.QueryRow(queryBestSeller, startDate, endDate).Scan(&report.ProdukTerlaris.Name, &report.ProdukTerlaris.QtyTerjual)
	if err != nil {
		if err == sql.ErrNoRows {
			report.ProdukTerlaris = models.BestSellingProduct{Name: "-", QtyTerjual: 0}
		} else {
			return report, err
		}
	}

	return report, nil
}
