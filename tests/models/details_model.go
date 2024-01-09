package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Details struct {
	Block       string `json:"block"`
	RoonType    string `json:"classroom"`
	Day         int    `json:"day"`
	HourSegment int    `json:"hours"`
}
type ColumnsData struct {
	Columns map[string]string
}

type Received struct {
	ID      primitive.ObjectID `bson:"_id"`
	RoomNo  string             `bson:"Room_no"`
	DayKey  int                `bson:"Day_key"`
	DayTime string             `bson:"Day/Time"`
	Columns ColumnsData        `bson:",inline"`
}
