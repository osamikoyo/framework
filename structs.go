package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
type DB struct{}
type Data struct{  
	url string
	db string
	collect string
}

var database Data = Data{
	url: "mongodb://localhost:27017",
	db: "OSamidb",
	collect: "animal",

}

func (d DB) add( document bson.M)  {
	// Set client options
	url := database.url
	collect := database.collect
	db := database.db
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)
		return
	}

	// Access the database
	database := client.Database(db)

	// Access the collection
	collection := database.Collection(collect)

	// Define a document
	

	// Insert the document
	result, err := collection.InsertOne(context.TODO(), document)
	if err != nil {
		fmt.Println("Error inserting document:", err)
		return
	}

	// Print the ID of the inserted document
	fmt.Println("Inserted document ID:", result.InsertedID)
}

type kytos struct{

  DB


}
