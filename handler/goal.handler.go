package handler

import (
	"fmt"
	"net/http"
	"time"

	model "github.com/SJ22032003/go-fullstack/model"
	util "github.com/SJ22032003/go-fullstack/util"
	goals_views "github.com/SJ22032003/go-fullstack/views/goals"
	gin "github.com/gin-gonic/gin"
)

const (
	GoalsCtxKey = "goals"
)

type GoalHandler struct{}

var goalsList = []model.Goal{}

func (g *GoalHandler) Goals(ctx *gin.Context) {
	goalsCtx := model.StoreGoalsToCtx(goalsList)
	template := util.NewTempl(goalsCtx, http.StatusOK, goals_views.GoalsPage())
	ctx.Render(http.StatusOK, template)
}

func (g *GoalHandler) AddGoal(ctx *gin.Context) {
	newGoal := model.NewGoal(
		ctx.PostForm("goal"),
		fmt.Sprintf("%d", time.Now().Nanosecond()),
		false,
	)
	goalsList = append(goalsList, newGoal)
	template := util.NewTempl(ctx, http.StatusOK, goals_views.GoalItem(newGoal))
	ctx.Render(http.StatusOK, template)
}

func (g *GoalHandler) DeleteGoal(ctx *gin.Context) {
	goalId := ctx.Param("id")
	for i, goal := range goalsList {
		if goal.Id == goalId {
			goalsList = append(goalsList[:i], goalsList[i+1:]...)
		}
	}
	ctx.String(http.StatusOK, "")
}
