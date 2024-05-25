package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/majoramari/visitor-logger/server"
	"github.com/majoramari/visitor-logger/utils"
)

func main() {
	router := server.NewRouter()

	// Enable CORS middleware
	http.Handle("/", utils.EnableCORS(router))

	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
