package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const db = "mydb"
const coll = "test"

var collection *mongo.Collection

type myDocument struct {
	MyMap   map[string]interface{} `json:"a"`
	MySlice []string               `json:"b"`
}

func init() {
	client, err := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal((err))
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	collection = client.Database(db).Collection(coll)
	fmt.Println("connection successfull")

}
func main() {
	fmt.Println("hii")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	doc := map[string]interface{}{
		"name":  "pi",
		"value": 3.14159,
	}
	res, err := collection.InsertOne(ctx, doc)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	id := res.InsertedID
	fmt.Printf("print id: %v", id)

}
