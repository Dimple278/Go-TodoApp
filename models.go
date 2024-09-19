package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var todoCollection *mongo.Collection

// ToDo represents a single to-do item
type ToDo struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
	CreatedAt time.Time          `bson:"createdAt"`
}

// Initialize the collection
func initCollection() {
	todoCollection = client.Database("todoApp").Collection("todos")
}

// AddToDo adds a new to-do item
func AddToDo(title string) error {
	newToDo := ToDo{
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	_, err := todoCollection.InsertOne(context.TODO(), newToDo)
	return err
}

// DeleteToDo deletes a to-do item by ID
func DeleteToDo(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = todoCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	return err
}

// ListToDos lists all to-do items
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

// MarkComplete marks a to-do item as complete
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

// MarkAllComplete marks all to-do items as complete
func MarkAllComplete() error {
	_, err := todoCollection.UpdateMany(
		context.TODO(),
		bson.M{},
		bson.D{{"$set", bson.D{{"completed", true}}}},
	)
	return err
}
