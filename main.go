package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/TedMartell/Blog_Aggregator/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	dbQueries := database.New(db) // Initialize dbQueries

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("PORT:", portString)

	cfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthz", handlerReadiness)
	mux.HandleFunc("/v1/err", handlerError)
	mux.HandleFunc("/v1/users", cfg.handlerCreateUser) // Use method with receiver

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
