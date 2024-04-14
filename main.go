package main

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)
var dataMONGO bson.M = DB.scan(DB{})
func main(){
	DATA := bson.M{
		"name" : "alice",
	}
	
    DB.add(DB{}, DATA)
	fmt.Println(dataMONGO)
	
}