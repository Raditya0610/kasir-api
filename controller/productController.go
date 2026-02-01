package controller

import (
	"kasir-api/models"
	"kasir-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *service.ProductService
}

func NewProductController(service *service.ProductService) *ProductController {
	return &ProductController{service: service}
}

// GetAllProducts godoc
// @Summary Ambil semua produk
// @Tags Products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func (h *ProductController) GetAllProducts(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID godoc
// @Summary Ambil detail satu produk
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func (h *ProductController) GetProductByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct godoc
// @Summary Tambah produk baru
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product Data"
// @Success 201 {object} models.Product
// @Router /products [post]
func (h *ProductController) CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Create(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, _ := h.service.GetByID(input.ID)
	c.JSON(http.StatusCreated, result)
}

// UpdateProduct godoc
// @Summary Update produk
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product Data"
// @Success 200 {object} models.Product
// @Router /products/{id} [put]
func (h *ProductController) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := h.service.Update(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Fetch ulang agar relasi category terbaru tampil
	result, _ := h.service.GetByID(updatedProduct.ID)
	c.JSON(http.StatusOK, result)
}

// DeleteProduct godoc
// @Summary Hapus produk
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Router /products/{id} [delete]
func (h *ProductController) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
