package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)
func main(){
	fmt.Println("hello world")
	DATA := bson.M{
		"name" : "alice",
	}
    DB.add(DB{}, DATA)
	
}