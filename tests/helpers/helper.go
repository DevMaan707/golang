package helpers

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

func Find(collection mongo.Collection) {
	found, err := collection.Find(context.Background())

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(found)
}
