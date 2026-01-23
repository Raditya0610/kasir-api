# 🏪 Kasir API

Kasir API adalah layanan RESTful API sederhana untuk sistem Point of Sales (POS). Proyek ini dibangun menggunakan bahasa pemrograman **Go (Golang)** dengan framework **Gin**, ORM **GORM**, dan database **PostgreSQL**.

Project ini dibuat untuk memenuhi tugas **Task Session 1** (Backend Development).

## 🚀 Teknologi yang Digunakan

* **Language:** [Go](https://go.dev/) (Golang)
* **Framework:** [Gin Web Framework](https://github.com/gin-gonic/gin)
* **Database:** PostgreSQL
* **ORM:** [GORM](https://gorm.io/)
* **Documentation:** [Swagger (Swaggo)](https://github.com/swaggo/swag)
* **Config:** [Godotenv](https://github.com/joho/godotenv)

## 📂 Struktur Project

```text
kasir-api/
├── config/         # Konfigurasi database
├── controllers/    # Logic handler untuk request API
├── docs/           # File generate Swagger documentation
├── models/         # Struct database (Schema)
├── routes/         # Definisi endpoint URL
├── .env            # Environment variables (buat .env anda sendiri)
├── main.go         # Entry point aplikasi
├── go.mod          # Dependency manager
└── README.md       # Dokumentasi project
```
