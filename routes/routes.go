package routes

import (
	"kasir-api/controller"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// --- Category Routes ---
	r.GET("/categories", controller.GetAllCategories)
	r.POST("/categories", controller.CreateCategory)
	r.GET("/categories/:id", controller.GetCategoryByID)
	r.PUT("/categories/:id", controller.UpdateCategory)
	r.DELETE("/categories/:id", controller.DeleteCategory)

	// --- Product Routes ---
	r.GET("/products", controller.GetAllProducts)
	r.POST("/products", controller.CreateProduct)
	r.GET("/products/:id", controller.GetProductByID)
	r.PUT("/products/:id", controller.UpdateProduct)
	r.DELETE("/products/:id", controller.DeleteProduct)

	return r
}
