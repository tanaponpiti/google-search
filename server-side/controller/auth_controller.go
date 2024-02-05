package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server-side/model"
	"server-side/response"
	"server-side/service"
	"strconv"
)

func Me(c *gin.Context) {
	userId, exist := c.Get("userId")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIdStr, ok := userId.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	userIdNum, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return
	}
	user, err := service.GetUserInfoFromId(userIdNum)
	complete := response.HandleErrorResponse(err, c)
	if complete {
		return
	}
	c.JSON(http.StatusOK, user)
	return
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func SignUp(c *gin.Context) {
	var userCreate model.UserCreate
	if err := c.ShouldBindJSON(&userCreate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	_, err := service.SignUp(userCreate)
	complete := response.HandleErrorResponse(err, c)
	if complete {
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sign up successfully"})
}

func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	username := req.Username
	password := req.Password
	token, err := service.Login(username, password)
	complete := response.HandleErrorResponse(err, c)
	if complete {
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		c.JSON(http.StatusBadRequest, gin.H{"error": "unauthorized"})
		return
	}
	token := authHeader[7:]

	err := service.Logout(token)
	complete := response.HandleErrorResponse(err, c)
	if complete {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
