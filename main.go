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
	db := InitDB()
	defer CloseDB()

	// Initialize Router
	r := mux.NewRouter()

	// Register Routes
	routes.RegisterRoutes(r, db)

	// Start Server
	port := ":8080"
	fmt.Println("🚀 Server started on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
