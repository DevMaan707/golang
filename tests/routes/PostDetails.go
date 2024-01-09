package routes

import (
	"fmt"

	helper "github.com/dev.maan707/golang/tests/helpers"
	"go.mongodb.org/mongo-driver/mongo"

	model "github.com/dev.maan707/golang/tests/models"
	"github.com/gin-gonic/gin"
)

const dbName = ""
const colName = ""

func HandleData(c *gin.Context, client *mongo.Client) {

	//creating instance of the structure
	var payload model.Details

	//binding the structure to the received data
	c.ShouldBindJSON(&payload)

	//Printing received
	fmt.Println("Data Successfully Received from the client")

	//Getting Collection

	collection := client.Database(dbName).Collection(colName)

	helper.Find(collection, payload.HourSegment)

}
