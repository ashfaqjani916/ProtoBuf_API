package controllers

import (
	"ProtoBuf_API/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func FetchTasks(c *gin.Context) {

	//connect to the google drive

	//fetch the tasks from the google drive file

	// decode the file to an array of tasks

	// return the array of tasks as response

	c.JSON(200, gin.H{
		"message": "FetchTasks",
	})
}
