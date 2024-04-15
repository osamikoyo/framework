package main

import (
	"fmt"


	"go.mongodb.org/mongo-driver/bson"
)

var dataMONGO []bson.M
func main(){
	DATA := bson.M{
		"name" : "alice",
	}
	
	dataMONGO = DB.scan(DB{}, bson.M{
		"name" : "alice",
	})
    DB.add(DB{}, DATA)
	fmt.Println(dataMONGO)
	
}