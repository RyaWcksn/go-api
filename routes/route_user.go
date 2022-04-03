package routes

import (
	"github.com/RyaWcksn/go-api/handler"
	"github.com/RyaWcksn/go-api/repositories"
	"github.com/RyaWcksn/go-api/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route gin.Engine) {
    userRepository := repositories.NewUserRepository(db)
    userService := services.NewUserService(*userRepository)
    userHandler := handler.NewUserHandler(*userService)


}
