package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_CONN")
	if dsn == "" {
		log.Fatal("DB_CONN is not set in environment variables")
	}

	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("🚀 Connected to Supabase (via pgx) successfully!")
}
