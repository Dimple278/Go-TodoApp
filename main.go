package main

import (
	"log"
	"net/http"

	"github.com/Dimple278/Go-TodoApp/controller"
	"github.com/Dimple278/Go-TodoApp/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func main() {
	models.ConnectMongoDB()
	models.InitCollection()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", controller.HomeHandler)
	http.HandleFunc("/todos", controller.AddToDoHandler)
	http.HandleFunc("/todos/", controller.DeleteToDoHandler)
	http.HandleFunc("/todos/complete/", controller.MarkCompleteHandler)
	http.HandleFunc("/todos/complete-all", controller.MarkAllCompleteHandler)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
