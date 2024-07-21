package controllers

import (
	"ProtoBuf_API/initializers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

type task struct {
	Id          string `json:"id"`
	Taskname    string `json:"taskname"`
	Description string `json:"description"`
	Due         string `json:"due"`
	Category    string `json:"category"`
}

func ModifyTasks(c *gin.Context) {

	var body struct {
		Tasks []task `json:"tasks"`
	}

	c.Bind(&body.Tasks)

	fmt.Println(body.Tasks)

	//call the fetch tasks function to get the tasks from the google drive file

	//modify the tasks array

	//encode the tasks array to a .proto file

	//save the .proto file to the google drive

	c.JSON(200, gin.H{
		"message": "ModifyTasks",
		"tasks":   body.Tasks,
	})
}
