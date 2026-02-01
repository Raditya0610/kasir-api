package repository

import (
	"database/sql"
	"errors"
	"kasir-api/models"
	"time"
)

type ProductRepository interface {
	FetchAll() ([]models.Product, error)
	FetchByID(id int) (models.Product, error)
	Store(product *models.Product) error
	Update(product *models.Product) error
	Delete(id int) error
}

type productRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FetchAll() ([]models.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, p.stock, p.category_id, p.created_at, p.updated_at,
		       c.id, c.name
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.deleted_at IS NULL
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		var c models.Category

		err := rows.Scan(
			&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.CreatedAt, &p.UpdatedAt,
			&c.ID, &c.Name,
		)
		if err != nil {
			return nil, err
		}
		p.Category = &c
		products = append(products, p)
	}
	return products, nil
}

func (r *productRepository) FetchByID(id int) (models.Product, error) {
	query := `
		SELECT p.id, p.name, p.price, p.stock, p.category_id, p.created_at, p.updated_at,
		       c.id, c.name
		FROM products p
		JOIN categories c ON p.category_id = c.id
		WHERE p.id = $1 AND p.deleted_at IS NULL
	`
	var p models.Product
	var c models.Category

	err := r.db.QueryRow(query, id).Scan(
		&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID, &p.CreatedAt, &p.UpdatedAt,
		&c.ID, &c.Name,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return p, errors.New("product not found")
		}
		return p, err
	}
	p.Category = &c
	return p, nil
}

func (r *productRepository) Store(p *models.Product) error {
	query := `
		INSERT INTO products (name, price, stock, category_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	now := time.Now()
	err := r.db.QueryRow(query, p.Name, p.Price, p.Stock, p.CategoryID, now, now).Scan(&p.ID)
	if err != nil {
		return err
	}
	p.CreatedAt = now
	p.UpdatedAt = now
	return nil
}

func (r *productRepository) Update(p *models.Product) error {
	query := `
		UPDATE products 
		SET name = $1, price = $2, stock = $3, category_id = $4, updated_at = $5
		WHERE id = $6 AND deleted_at IS NULL
	`
	p.UpdatedAt = time.Now()
	res, err := r.db.Exec(query, p.Name, p.Price, p.Stock, p.CategoryID, p.UpdatedAt, p.ID)
	if err != nil {
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("product not found or no change")
	}
	return nil
}

func (r *productRepository) Delete(id int) error {
	query := `UPDATE products SET deleted_at = $1 WHERE id = $2`
	_, err := r.db.Exec(query, time.Now(), id)
	return err
}
