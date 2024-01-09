package controllers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = ""

func ConnectDB() (*mongo.Client, error) {
	//client options
	clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongoDB

	client, err := mongo.Connect(context.TODO(), clientOption)

	//Handling Errors which may raise
	if err != nil {
		log.Fatal((err))
		return nil, err
	}

	fmt.Println("Successful Connection with MongoDB")

	return client, nil
}
