package handlers

import (
	"brainyquiz/internal/domain/entities"
	"brainyquiz/internal/domain/services"

	"github.com/gin-gonic/gin"
)

type FriendHandler struct {
	friendService services.FriendService
}

func NewFriendHandler(friendService services.FriendService) *FriendHandler {
	return &FriendHandler{friendService: friendService}
}

func (h *FriendHandler) AddFriend(c *gin.Context) {
	var friend entities.Friend
	if err := c.ShouldBindJSON(&friend); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.friendService.AddFriend(c, friend); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Friend added successfully"})
}

func (h *FriendHandler) RemoveFriend(c *gin.Context) {
	var friend entities.Friend

	if err := c.ShouldBindJSON(&friend); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := h.friendService.RemoveFriend(c, friend); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Friend removed successfully"})

}

func (h *FriendHandler) GetFriends(c *gin.Context) {
	username := c.Param("username")
	friends, err := h.friendService.GetFriends(c, username)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, friends)
}
