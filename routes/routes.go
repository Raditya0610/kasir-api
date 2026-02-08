package routes

import (
	"kasir-api/controller"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(productCtrl *controller.ProductController, categoryCtrl *controller.CategoryController, transactionCtrl *controller.TransactionController) *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())
	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/swagger/index.html")
	})

	// --- Category Routes ---
	r.GET("/categories", categoryCtrl.GetAllCategories)
	r.POST("/categories", categoryCtrl.CreateCategory)
	r.GET("/categories/:id", categoryCtrl.GetCategoryByID)
	r.PUT("/categories/:id", categoryCtrl.UpdateCategory)
	r.DELETE("/categories/:id", categoryCtrl.DeleteCategory)

	// --- Product Routes ---
	r.GET("/products", productCtrl.GetAllProducts)
	r.POST("/products", productCtrl.CreateProduct)
	r.GET("/products/:id", productCtrl.GetProductByID)
	r.PUT("/products/:id", productCtrl.UpdateProduct)
	r.DELETE("/products/:id", productCtrl.DeleteProduct)

	// --- Transaction Routes ---
	r.POST("/checkout", transactionCtrl.HandleCheckout)

	return r
}
