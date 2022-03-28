package main

import (
	"fmt"

	"github.com/RyaWcksn/go-api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	//Initialize Database
	db := config.Connection()
	fmt.Println(db)

	// Initialize Gin Gonic
	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.Run(":6969")
}
