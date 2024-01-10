package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//Struct defining the data which will be received by the backend from app
type Details struct {
	Block       string `json:"block"`
	RoonType    string `json:"classroom"`
	Day         int    `json:"day"`
	HourSegment int    `json:"hours"`
}

type ColumnsData struct {
	Columns map[string]string
}

//Struct defining the data which will be received by the backend from mongoDB

type Received struct {
	ID      primitive.ObjectID `bson:"_id"`
	RoomNo  string             `bson:"Room_no"`
	DayKey  int                `bson:"Day_key"`
	DayTime string             `bson:"Day/Time"`
	Columns ColumnsData        `bson:",inline"`
}

type Reserve struct {
	Room_No string `bson:"Room_No"`
	Hour    int    `bson:"Hour"`
}
