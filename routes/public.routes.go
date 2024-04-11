package routes

import (
	"github.com/SJ22032003/go-fullstack/handler"
	gin "github.com/gin-gonic/gin"
)

func PublicRoutes(server *gin.Engine, v *gin.RouterGroup) {

	goalHandler := handler.GoalHandler{}
	favouriteHandler := handler.FavouriteHandler{}

	goalsRoute := v.Group("/public/goals")
	{
		goalsRoute.GET("/", goalHandler.Goals)
		goalsRoute.POST("/add-goal", goalHandler.AddGoal)
		goalsRoute.DELETE("/delete-goal/:id", goalHandler.DeleteGoal)
	}

	favouriteRoute := v.Group("/public/favourites")
	{
		favouriteRoute.GET("/", favouriteHandler.Favourites)
	}

}
