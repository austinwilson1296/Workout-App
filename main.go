package main

import (
    "log"
    "net/http"
    "os"
    "database/sql"
    // "embed"
    "github.com/austinwilson1296/fitted/routes"
    "github.com/austinwilson1296/fitted/internal/database"
    "github.com/austinwilson1296/fitted/config"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

// var staticFS embed.FS

func main() {
    const port = "8080"

    if os.Getenv("ENVIRONMENT") != "production" {
        if err := godotenv.Load(); err != nil {
            log.Fatal("Error loading .env file")
        }
    }
    

    dbURL := os.Getenv("DB_URL")
    if dbURL == "" {
        log.Fatal("DB_URL must be set")
    }

    platform := os.Getenv("PLATFORM")
    if platform == "" {
        log.Fatal("PLATFORM must be set")
    }
    jwtSecretKey := os.Getenv("jwtSecret")
	if jwtSecretKey == "" {
		log.Fatal("JWT_SECRET environment variable is not set")
	}

    dbConn, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatalf("Error opening database: %s", err)
    }
    dbQueries := database.New(dbConn)

    apiConfig := &config.ApiCfg{
        DB:        dbQueries,
        Platform:  platform,
        JwtSecret: jwtSecretKey,
    }


    mux := http.NewServeMux()
    routes.RegisterRoutes(mux,apiConfig)
    

    srv := &http.Server{
        Addr:    ":" + port,
        Handler: mux,
    }

    log.Printf("Serving on port: %s\n", port)
    log.Fatal(srv.ListenAndServe())
}

