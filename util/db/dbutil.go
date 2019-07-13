package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Util provides Mongo db related functions.
type Util struct {
}

func connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	return client
}

// Ping sends a ping to the client, and returns true if there are no errors.
func (d *Util) Ping() bool {
	client := connect()
	defer disconnect(client)
	err := client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func disconnect(client *mongo.Client) bool {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func getTable(client *mongo.Client, database, collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
