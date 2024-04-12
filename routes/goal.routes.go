package routes

import (
	"github.com/SJ22032003/go-fullstack/handler"
	gin "github.com/gin-gonic/gin"
)

func GoalRoutes(v *gin.RouterGroup) {

	goalHandler := handler.GoalHandler{}

	goalsRoute := v.Group("/goals")
	{
		goalsRoute.GET("/", goalHandler.Goals)
		goalsRoute.POST("/add-goal", goalHandler.AddGoal)
		goalsRoute.DELETE("/delete-goal/:id", goalHandler.DeleteGoal)
	}

}
