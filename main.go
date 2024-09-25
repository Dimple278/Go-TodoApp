package main

import (
	"log"
	"net/http"

	"github.com/Dimple278/Go-TodoApp/models"
	"github.com/Dimple278/Go-TodoApp/router"
)

func main() {
	// Connect to MongoDB and initialize the collection
	models.ConnectMongoDB()
	models.InitCollection()

	// Setup the routes
	router.SetupRoutes()

	// Start the server
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
