package routes

import "github.com/gin-gonic/gin"

func PublicRoutes(v *gin.RouterGroup) {
	publicRoute := v.Group("/public")
	{
		publicRoute.GET("/", func(c *gin.Context) {})
	}
}
