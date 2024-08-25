package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviroment")
	}

	fmt.Println("PORT:", portString)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /api/err", handlerError)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())

}
