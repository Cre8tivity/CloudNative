package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var coll *mongo.Collection

	// Find all documents in which the "name" field is "Bob".
	// Specify the Sort option to sort the returned documents by age in
	// ascending order.
	opts := options.Find().SetSort(bson.D{{"age", 1}})
	cursor, err := coll.Find(context.TODO(), bson.D{{"name", "Bob"}}, opts)
	if err != nil {
		log.Fatal(err)
	}

	// Get a list of all returned documents and print them out.
	// See the mongo.Cursor documentation for more examples of using cursors.
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
