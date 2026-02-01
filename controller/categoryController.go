package controller

import (
	"kasir-api/models"
	"kasir-api/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// CreateCategory godoc
// @Summary Tambah kategori baru
// @Tags Categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category Data"
// @Success 201 {object} models.Category
// @Router /categories [post]
func (h *CategoryController) CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.Create(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

// GetAllCategories godoc
// @Summary Ambil semua kategori
// @Tags Categories
// @Produce json
// @Success 200 {array} models.Category
// @Router /categories [get]
func (h *CategoryController) GetAllCategories(c *gin.Context) {
	categories, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetCategoryByID godoc
// @Summary Ambil detail satu kategori
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} models.Category
// @Router /categories/{id} [get]
func (h *CategoryController) GetCategoryByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory godoc
// @Summary Update kategori
// @Tags Categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category Data"
// @Success 200 {object} models.Category
// @Router /categories/{id} [put]
func (h *CategoryController) UpdateCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var input models.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := h.service.Update(id, input)
	if err != nil {
		// Bisa improve error handling untuk membedakan 404 dan 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory godoc
// @Summary Hapus kategori
// @Tags Categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Router /categories/{id} [delete]
func (h *CategoryController) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}
