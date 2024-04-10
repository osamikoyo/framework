package main

import (

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

)

type Mongo interface{
	add(collect string)
    del(filter string)
	scan(filter string)
}
