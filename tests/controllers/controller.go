package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""
const dbName = ""
const colName = ""

var collection *mongo.Collection

func init() {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB

	client, err := mongo.Connect(context.TODO(), clientOption)

	//Handling Errors which may raise
	if err != nil {
		log.Fatal((err))
	}

	fmt.Println("Successful Connection with MongoDB")

	//Getting the Collection
	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection instance ready")

}
