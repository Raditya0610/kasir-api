package service

import (
	"kasir-api/models"
	"kasir-api/repository"
	"time"
)

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}

func (s *TransactionService) GetDailyReport() (models.SalesReport, error) {
	now := time.Now()
	startDate := now.Format("2006-01-02") + " 00:00:00"
	endDate := now.Format("2006-01-02") + " 23:59:59"

	return s.repo.GetSalesReport(startDate, endDate)
}
