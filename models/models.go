package models

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var todoCollection *mongo.Collection
var client *mongo.Client

type ToDo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
	CreatedAt time.Time          `bson:"createdAt"`
}

func ConnectMongoDB() {

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

func InitCollection() {
	todoCollection = client.Database("todoApp").Collection("todos")
}

func AddToDo(title string) error {
	newToDo := ToDo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	_, err := todoCollection.InsertOne(context.TODO(), newToDo)
	return err
}

func DeleteToDo(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = todoCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

func ListToDos() ([]ToDo, error) {
	var todos []ToDo

	cursor, err := todoCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var todo ToDo
		err := cursor.Decode(&todo)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func MarkComplete(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = todoCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objID},
		bson.D{{"$set", bson.D{{"completed", true}}}},
	)
	return err
}

func MarkAllComplete() error {
	_, err := todoCollection.UpdateMany(
		context.TODO(),
		bson.M{},
		bson.D{{"$set", bson.D{{"completed", true}}}},
	)
	return err
}
