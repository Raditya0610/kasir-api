package controller

import (
	"kasir-api/models"
	"kasir-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service *service.TransactionService
}

func NewTransactionController(service *service.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

// Checkout godoc
// @Summary Checkout products
// @Tags Transactions
// @Accept json
// @Produce json
// @Param checkout body models.CheckoutRequest true "Checkout Data"
// @Success 200 {object} models.Transaction
// @Router /checkout [post]
func (h *TransactionController) HandleCheckout(c *gin.Context) {
	var req models.CheckoutRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// GetDailyReport godoc
// @Summary Get sales report for today
// @Tags Reports
// @Produce json
// @Success 200 {object} models.SalesReport
// @Router /report/hari-ini [get]
func (h *TransactionController) GetDailyReport(c *gin.Context) {
	report, err := h.service.GetDailyReport()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}
