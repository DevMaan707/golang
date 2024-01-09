package helpers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/dev.maan707/golang/tests/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(collection *mongo.Collection, hour int) {

	//Getting todays Day
	currentTime := time.Now()

	day := currentTime.Weekday()

	//Defining my filter criteria

	filter := bson.D{
		{"Day_Key", day - 1},
	}

	//filter := bson.D{
	//	{"Day/Time", "Tuesday"},
	//	{"Columns.6", bson.D{
	//		{"$regex", "TRAINING$|LAB$"},
	//	}},
	//}

	fmt.Println("Initiating Filter")
	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Got cursor , Searching values: ")

	defer cursor.Close(context.Background())

	//Iterating through the results

	for cursor.Next(context.Background()) {
		var data models.Received
		err := cursor.Decode(&data)

		if err != nil {
			log.Fatal(err)
		}
		var rooms []string
		for colname, value := range data.Columns {

			if colname == hour {
				if strings.HasSuffix(value, "TRAINING") || strings.HasSuffix(value, "LAB") {
					fmt.Printf("Match Found : Room Number - %s\n", data.RoomNo)
					rooms = append(rooms, data.RoomNo)
				}
			}
		}
	}

	//fmt.Println(found)
}
