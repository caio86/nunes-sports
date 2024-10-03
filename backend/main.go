package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caio86/nunes-sports/backend/internal/adapters/input/api/handlers/product"
	"github.com/caio86/nunes-sports/backend/internal/adapters/output/database"
	"github.com/caio86/nunes-sports/backend/internal/core/service"
	"github.com/caio86/nunes-sports/backend/internal/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	dbname := "db"
	password := "123"

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := database.NewProductRepository(db)
	svc := service.NewProductService(repo)
	handler := product.New(svc)

	router := http.NewServeMux()
	router.HandleFunc("GET /product", handler.Get)
	router.HandleFunc("GET /product/{id}", handler.GetByID)
	router.HandleFunc("POST /product", handler.Create)

	apiVersion := http.NewServeMux()
	apiVersion.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(apiVersion),
	}

	log.Println("Starting server on port :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
