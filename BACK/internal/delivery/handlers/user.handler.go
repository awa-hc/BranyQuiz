package handlers

import (
	"brainyquiz/internal/domain/entities"
	"brainyquiz/internal/domain/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}
func (h *UserHandler) CreateUser(c *gin.Context) {
	var User entities.User
	if err := c.ShouldBindJSON(&User); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.CreateUser(c, &User); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "User created successfully"})
}
