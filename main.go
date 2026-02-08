package main

import (
	"kasir-api/config"
	"kasir-api/controller"
	"kasir-api/docs"
	"kasir-api/repository"
	"kasir-api/routes"
	"kasir-api/service"
	"log"
	"os"

	_ "kasir-api/docs"

	_ "github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// @title           Kasir API
// @version         1.0
// @description     API Server untuk aplikasi Kasir sederhana
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      kasir-api-production.up.railway.app
// @BasePath  /
// @schemes   https

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Using system environment variables.")
	}
	config.ConnectDatabase()

	defer config.DB.Close()

	// --- Category Layer ---
	categoryRepo := repository.NewCategoryRepository(config.DB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryCtrl := controller.NewCategoryController(categoryService)

	// --- Product Layer ---
	productRepo := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepo)
	productCtrl := controller.NewProductController(productService)

	// --- Transaction Layer ---
	transactionRepo := repository.NewTransactionRepository(config.DB)
	transactionService := service.NewTransactionService(transactionRepo)
	transactionCtrl := controller.NewTransactionController(transactionService)

	r := routes.SetupRouter(productCtrl, categoryCtrl, transactionCtrl)

	// 4. Run Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port lokal
	}

	railwayDomain := os.Getenv("RAILWAY_PUBLIC_DOMAIN")

	if railwayDomain != "" {
		docs.SwaggerInfo.Host = railwayDomain
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		docs.SwaggerInfo.Host = "localhost:" + port
		docs.SwaggerInfo.Schemes = []string{"http"}
	}
	r.Run(":" + port)
}
