package datastore

import (
	"context"
	"fmt"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* Initialized and exposed through  GetMongoClient().*/
var clientInstance *mongo.Client

//Used during creation of singleton client object in GetMongoClient().
var clientInstanceError error

//Used to execute client creation procedure only once.
var mongoOnce sync.Once

//I have used below constants just to hold required database config's.
const (
	CONNECTIONSTRING = "mongodb://localhost:27017"
	DB               = "Chatbot"
	ISSUES           = "profileActivity"
)

//GetMongoClient - Return mongodb connection to work with
func GetMongoClient() (*mongo.Client, error) {
	//Perform connection creation operation only once.
	mongoOnce.Do(func() {
		// Set client options
		clientOptions := options.Client().ApplyURI(CONNECTIONSTRING)
		// Connect to MongoDB
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		// Check the connection
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}

func UserActivityInsert() {

	//Get MongoDB connection using connectionhelper.
	client, err := GetMongoClient()
	if err != nil {
		return
	}
	//Create a handle to the respective collection in the database.
	collection := client.Database(DB).Collection("profileActivity")
	//Perform InsertOne operation & validate against the error.
	res, err := collection.InsertOne(context.TODO(), bson.D{{"name", "FirstNM"}, {"value", "SenderTxt"}})
	if err != nil {
		return
	}
	//Return success without any error.
	id := res.InsertedID

	fmt.Printf("Document inserted succesfully: %+v\n", id)

}
