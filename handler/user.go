package handler

import (
	"net/http"

	"github.com/RyaWcksn/go-api/entities"
	"github.com/RyaWcksn/go-api/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
    service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
    return &UserHandler{service}
}

func (h UserHandler) Create(c *gin.Context) {
    var model entities.User

    if err := c.ShouldBindJSON(&model); err != nil {
        panic(err.Error())
    }
    user, err := h.service.Create(model)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "data": nil,
            "error": err.Error(),
        })
    }
    c.JSON(http.StatusOK, gin.H{
        "data": user,
        "error": nil,
    })
}
