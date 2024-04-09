package main

import (
"context"
"log"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
	"fmt"
)

type Mongo interface{
	add(collect string)
    del(filter string)
	scan(filter string)
}