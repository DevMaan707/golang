package routes

import (
	"fmt"
	"net/http"

	helper "github.com/dev.maan707/golang/tests/helpers"
	"go.mongodb.org/mongo-driver/mongo"

	model "github.com/dev.maan707/golang/tests/models"
	"github.com/gin-gonic/gin"
)

const dbName = "TimeTable"
const colName = "A_Block"

func HandleData(c *gin.Context, client *mongo.Client) {

	//creating instance of the structure
	var payload model.Details

	//binding the structure to the received data
	c.ShouldBindJSON(&payload)

	//Printing received
	fmt.Println("Data Successfully Received from the client")

	//Getting Collection

	collection := client.Database(dbName).Collection(colName)

	rooms := helper.Find(collection, payload.HourSegment, payload.Block, payload.Day)
	var length = 5
	if len(rooms) < 5 {
		length = len(rooms) - 1
	}
	response := map[string]interface{}{
		"number":    length,
		"classroom": rooms,
	}
	c.JSON(http.StatusOK, response)

	fmt.Println("Response Sent!")
}
