package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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

func (d DB) add( document bson.M) {
	// Set client options
	url := database.url
	collect := database.collect
	db := database.db
	clientOptions := options.Client().ApplyURI(url)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)

	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("Error pinging MongoDB:", err)

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
	}

	// Print the ID of the inserted document
	fmt.Println("Inserted document ID:", result.InsertedID)
}


func (d DB) scan(filter bson.M) ([]bson.M){
	    // Set client options
		clientOptions := options.Client().ApplyURI(database.url)

		// Connect to MongoDB
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
	
		// Check the connection
		err = client.Ping(context.Background(), nil)
		if err != nil {
			log.Fatal(err)
		}
	
		// Get a handlonto the "users" collection in the "mydb" database
		collection := client.Database(database.db).Collection(database.collect)
	
		// Define a filter to match documents with a specific email address
	
		// Find the user document(s) matching the filter
		var results []bson.M
		cursor, err := collection.Find(context.Background(), filter)
		if err != nil {
			log.Fatal(err)
		}
	
		// Iterate over the cursor and decode the documents into User structs
		defer cursor.Close(context.Background())
		for cursor.Next(context.Background()) {
			var user bson.M
			err := cursor.Decode(&user)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, user)
		}
	
		fmt.Printf("Found %d user(s):\n", len(results))
		
		return results
	
	
}
