package main

import (
	"fmt"

	"ProtoBuf_API/controllers"
	"ProtoBuf_API/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello World")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API is running successfully",
		})
	})

	r.POST("/fetchTasks", controllers.FetchTasks)
	r.POST("/modifyTasks", controllers.ModifyTasks)
	r.Run()
}
