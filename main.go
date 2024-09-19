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
	// Define routes
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/add", AddToDoHandler)
	http.HandleFunc("/delete", DeleteToDoHandler)
	http.HandleFunc("/complete", MarkCompleteHandler)
	http.HandleFunc("/complete-all", MarkAllCompleteHandler)

	log.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
