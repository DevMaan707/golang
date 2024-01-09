package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Details struct {
	RoonType    string `json:"classroom"`
	HourSegment int    `json:"hour"`
}

type Received struct {
	ID      primitive.ObjectID `bson:"_id omitempty"`
	RoomNo  string             `json:"Room_no"`
	DayKey  int                `bson:"Day_key"`
	DayTime string             `bson:"Day/Time"`
	Columns map[int]string
}
