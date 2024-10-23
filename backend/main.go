package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/caio86/nunes-sports/backend/internal/adapters/input/api/handlers/product"
	"github.com/caio86/nunes-sports/backend/internal/adapters/output/database"
	"github.com/caio86/nunes-sports/backend/internal/core/service"
	"github.com/caio86/nunes-sports/backend/internal/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASS")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	repo := database.NewProductRepository(conn)
	svc := service.NewProductService(repo)
	handler := product.New(svc)

	router := http.NewServeMux()
	router.HandleFunc("GET /product", handler.Get)
	router.HandleFunc("GET /product/{id}", handler.GetByID)
	router.HandleFunc("POST /product", handler.Create)
	router.HandleFunc("PUT /product/{id}", handler.Update)
	router.HandleFunc("DELETE /product/{id}", handler.Delete)

	apiVersion := http.NewServeMux()
	apiVersion.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	dist := "../frontend/dist/acme-shop"
	files := http.FileServer(http.Dir(dist))
	apiVersion.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Try to serve the file directly
		_, err := http.Dir(dist).Open(r.URL.Path)
		if err != nil {
			// If the file doesn't exists, serve index.html
			http.ServeFile(w, r, dist+"/index.html")
		} else {
			// If the file exists, serve it
			files.ServeHTTP(w, r)
		}
	})

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.CORS,
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
