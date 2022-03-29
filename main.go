package main

import (
	"github.com/RyaWcksn/go-api/candidate"
	"github.com/RyaWcksn/go-api/config"
	"github.com/RyaWcksn/go-api/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	//Initialize Database
	db := config.Connection()
	candidatesRepository := candidate.NewCandidateRepository(db)
	candidateService := candidate.NewCandidateService(candidatesRepository)
	candidateHandler := handler.NewCandidateHandler(candidateService)

	route := gin.Default()
	route.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	route.POST("/candidates", candidateHandler.CandidatePostHandler)
	route.Run(":6969")
}
