package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/starlingvibes/Simple-CRUD-MonGo/models"
)

type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required"`
}

// register user with username and password
func RegisterUser(c *gin.Context) {
    var req models.User
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
}
