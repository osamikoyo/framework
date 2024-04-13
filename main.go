package main

import (
	"go.mongodb.org/mongo-driver/bson"
)
func main(){
	DATA := bson.M{
		"name" : "alice",
	}
    DB.add(DB{}, DATA)
}