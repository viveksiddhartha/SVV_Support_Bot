package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Datastore interface {
}

type RDatastore interface {
}

//var db sql.DB

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, "svcrm:Pass#word1@tcp(localhost:3306)/sv_crm")
	db.SetConnMaxLifetime(20)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(5)
	db.Stats()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err.Error())
	}

	return db

}

func MongoConn() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

}

/*
func RDBConn() (dbs *sql.DB) {
	dbDriver := "redis"
	dbs, err := sql.Open(dbDriver, "sv_crm:sv_crm@tcp(127.0.0.1:3306)/SV_CRM")
	if err != nil {
		panic(err.Error())
	}
	return dbs

}
*/
