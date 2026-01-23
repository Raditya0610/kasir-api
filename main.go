package main

import (
	"kasir-api/config"
	_ "kasir-api/docs"
	"kasir-api/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title Kasir API
// @version 1.0
// @description Ini adalah API server untuk Kasir App tugas session 1.
// @BasePath /
func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("env not loaded")
	}
	config.ConnectDatabase()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
