package controller

import (
	"kasir-api/config"
	"kasir-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct godoc
// @Summary Tambah produk baru
// @Description Menambahkan data produk ke database dengan relasi kategori
// @Tags Products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product Data"
// @Success 201 {object} models.Product
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cek apakah CategoryID valid (Opsional, tapi praktik bagus)
	var category models.Category
	if err := config.DB.First(&category, product.CategoryID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID not found"})
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Load data kategori agar response lengkap
	config.DB.Preload("Category").First(&product, product.ID)

	c.JSON(http.StatusCreated, product)
}

// GetAllProducts godoc
// @Summary Ambil semua produk
// @Description Mengambil list semua produk beserta kategorinya
// @Tags Products
// @Produce json
// @Success 200 {array} models.Product
// @Router /products [get]
func GetAllProducts(c *gin.Context) {
	var products []models.Product
	// Preload("Category") berfungsi seperti JOIN untuk mengambil data relasi
	config.DB.Preload("Category").Find(&products)
	c.JSON(http.StatusOK, products)
}

// GetProductByID godoc
// @Summary Ambil detail satu produk
// @Description Mengambil data produk berdasarkan ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	var product models.Product
	if err := config.DB.Preload("Category").First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// UpdateProduct godoc
// @Summary Update produk
// @Description Mengubah data produk berdasarkan ID
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.Product true "Product Data"
// @Success 200 {object} models.Product
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Jika CategoryID diubah, validasi lagi
	if input.CategoryID != 0 {
		var category models.Category
		if err := config.DB.First(&category, input.CategoryID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category ID not found"})
			return
		}
	}

	config.DB.Model(&product).Updates(input)

	// Reload data untuk menampilkan hasil update beserta relasi
	config.DB.Preload("Category").First(&product, product.ID)

	c.JSON(http.StatusOK, product)
}

// DeleteProduct godoc
// @Summary Hapus produk
// @Description Menghapus (soft delete) produk berdasarkan ID
// @Tags Products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.First(&product, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	config.DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
