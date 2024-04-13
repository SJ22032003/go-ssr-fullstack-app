package routes

import (
	"github.com/SJ22032003/go-fullstack/handler"
	"github.com/gin-gonic/gin"
)

func PublicRoutes(v *gin.RouterGroup) {

	authHandler := handler.AuthHandler{}

	publicRoute := v.Group("/")
	{
		publicRoute.GET("/", authHandler.AuthPage)
		publicRoute.POST("/validate", authHandler.InputValidation)
	}
}
