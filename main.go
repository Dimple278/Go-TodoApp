package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func connectMongoDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	uri := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(uri)

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}

func main() {
	connectMongoDB()
	initCollection()

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/todos", AddToDoHandler)
	http.HandleFunc("/todos/", DeleteToDoHandler)
	http.HandleFunc("/todos/complete/", MarkCompleteHandler)
	http.HandleFunc("/todos/complete-all", MarkAllCompleteHandler)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
