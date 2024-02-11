package route

import (
	"github.com/gin-gonic/gin"
	"server-side/controller"
	"server-side/service"
)

func RegisterAuthRoutes(rg *gin.RouterGroup) {
	authGroup := rg.Group("/auth")
	authGroup.GET("/me", service.AuthMiddleware(), controller.Me)
	authGroup.POST("/login", controller.Login)
	authGroup.POST("/logout", controller.Logout)
	authGroup.POST("/signup", controller.SignUp)
}
