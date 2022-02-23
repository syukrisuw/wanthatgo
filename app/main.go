package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Error: %v \n", err)
		}

	}()

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run("127.0.0.1:9090")

}
