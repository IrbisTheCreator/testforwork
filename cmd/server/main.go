package main

import (
	"flag"
	"fmt"
	"irbistest/handlers"
	"irbistest/iternal/app"
	"irbistest/routes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	var cfg app.Config
	flag.StringVar(&cfg.DB.Host, "db-host", getEnvOrDefault("DB_HOST", "localhost"), "PostgreSQL host")
	flag.StringVar(&cfg.DB.Port, "db-port", getEnvOrDefault("DB_PORT", "5432"), "PostgreSQL port")
	flag.StringVar(&cfg.DB.User, "db-user", getEnvOrDefault("DB_USER", "postgres"), "PostgreSQL user")
	flag.StringVar(&cfg.DB.Password, "db-password", getEnvOrDefault("DB_PASSWORD", "200670"), "PostgreSQL password")
	flag.StringVar(&cfg.DB.Name, "db-name", getEnvOrDefault("DB_NAME", "testdb"), "PostgreSQL database name")
	flag.StringVar(&cfg.Server.Port, "port", getEnvOrDefault("PORT", "8080"), "Server port")
	flag.Parse()

	application := app.NewApplication(cfg)

	if err := application.ConnectToDatabase(); err != nil {
		log.Fatal("Could not connect to PostgreSQL: ", err)
	}

	handlers := handlers.NewHandlers(application)

	router := routes.SetupRouter(handlers)

	port := cfg.Server.Port
	fmt.Printf("Server started at http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
