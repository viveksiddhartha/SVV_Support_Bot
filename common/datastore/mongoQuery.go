package datastore

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func userActivityInsert() {

	collection := client.Database("Chatbot").Collection("profileActivity")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.D{{"name", "pi"}, {"value", 3.14159}})
	if err != nil {
		log.Fatal(err)
	}
	id := res.InsertedID

	fmt.Printf("Document inserted succesfully: %+v\n", id)
}
