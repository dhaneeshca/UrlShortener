package main

import (
	"UrlShortener/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize Database
	InitDB()
	defer CloseDB()

	// Initialize Router
	r := mux.NewRouter()

	// Register Routes
	routes.RegisterRoutes(r)

	// Start Server
	port := ":8080"
	fmt.Println("🚀 Server started on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
