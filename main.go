package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	engine := gin.Default()

	engine.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 200,
			"payload": "hello gin",
		})
	})

	fmt.Println(">>> Gin Start <<<")
	engine.Run("localhost:3000")
}