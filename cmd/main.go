package main

import (
	"github.com/SJ22032003/go-fullstack/routes"
	gin "github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.Static("/static", "./static")

	v := server.Group("/")

	routes.GoalRoutes(v)
	routes.PublicRoutes(v)

	server.Run(":8080")

}
