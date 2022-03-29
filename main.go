package main

import (
	"github.com/RyaWcksn/go-api/config"
	"github.com/RyaWcksn/go-api/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Routing
	router := SetupRouter()

	router.Run(":6969")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	db := config.Connection()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	routes.InitCandidateRoute(db, router)

	return router
}
