package main

import (
"context"
"fmt"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)
type Mongo interface{
	add(collect string)
    del(filter string)
	scan(filter string)
}

type Animal struct{
  age int
  name string
}
func (a Animal) add(url string){
   // Set client options
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
database := client.Database("myDatabase")

// Access the collection
collection := database.Collection("myCollection")

// Define a document
document := bson.M{
	"key": "value",
}

// Insert the document
result, err := collection.InsertOne(context.TODO(), document)
if err != nil {
	fmt.Println("Error inserting document:", err)
	return
}

// Print the ID of the inserted document
fmt.Println("Inserted document ID:", result.InsertedID)
}

