package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"database/sql"
	"github.com/austinwilson1296/fitted/handlers"  
	"github.com/austinwilson1296/fitted/internal/database"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

type ApiCfg struct {
	platform string
	jwtSecret string
	db        *database.Queries  
}

func main() {
	const filepathRoot = "./templates"
	const port = "8080"

	
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}

	
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}

	// Open database connection
	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(dbConn)

	
	apiConfig := &ApiCfg{
		db:        dbQueries,
		platform:  platform,
	}

	
	mux := http.NewServeMux()

	
	fmt.Printf("API Config: %+v\n", apiConfig)

	
	mux.HandleFunc("/", handlers.HandlerHomePage)
	mux.HandleFunc("/workout", handlers.HandlerWorkOutPage)
	mux.HandleFunc("/api/warmup", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerGenerateWarmUp(apiConfig, w, r)
	})

	
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	
	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

