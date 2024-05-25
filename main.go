package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/majoramari/visitor-logger/server"
	"github.com/majoramari/visitor-logger/utils"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("No .env file")
	}

	router := server.NewRouter()

	// Enable CORS middleware
	http.Handle("/", utils.EnableCORS(router))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
