package helpers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/dev.maan707/golang/tests/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Find(collection *mongo.Collection, hour int, Block string, Day int) []string {

	fmt.Printf("HourSegment = %d\nBlock = %s\nDay= %d\n", hour, Block, Day)

	if err := collection.Database().Client().Ping(context.Background(), nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	//Describing the filter

	filter := bson.M{
		strconv.Itoa(hour): bson.M{"$regex": "(TRAINING|LAB)$"},
		"Day_Key":          Day,
	}

	//Initiating the Find Operation
	fmt.Println("Initiating Filter")
	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("Got cursor , Searching values: ")

	defer cursor.Close(context.Background())

	//Checking the length of the cursor
	fmt.Println("Cursor Count:", cursor.RemainingBatchLength())

	//Iterating through the results

	var rooms []string
	var data models.Received

	//Trying to iterate through the received data
	for cursor.Next(context.Background()) {

		fmt.Println("Decoding the json")
		err := cursor.Decode(&data)

		if err != nil {
			log.Fatal(err)
		}

		//printing the object ID just incase
		fmt.Println(data.ID)

		//This code was to manually search the data

		/* for colname, value := range data.Columns.Columns {
			if colname == strconv.Itoa(hour) {
				fmt.Println("Checking column == hour.....")
				if strings.HasSuffix(value, "TRAINING") || strings.HasSuffix(value, "LAB") {
					fmt.Printf("Match Found : Room Number - %s\n", data.RoomNo)
					rooms = append(rooms, data.RoomNo)
				}
			}
		} */

		//appening the value of rooms to a new slice
		rooms = append(rooms, data.RoomNo)

	}

	//Iterating through the rooms just incase
	for _, room := range rooms {
		fmt.Println(room)
	}

	//Returning the slice back to routes/PostDetails.go
	return rooms

}
