package model

import "context"

type keyType string

const (
	GoalsCtxKey keyType = "goals"
)

type Goal struct {
	Name      string
	Id        string
	Favourite bool
}

func NewGoal(name string, id string, favourite bool) Goal {
	return Goal{
		Name:      name,
		Id:        id,
		Favourite: favourite,
	}
}

func StoreGoalsToCtx(goalsList []Goal) context.Context {
	return context.WithValue(context.Background(), GoalsCtxKey, goalsList)
}

func GetGoalsFromCtx(ctx context.Context) []Goal {
	if goals, ok := ctx.Value(GoalsCtxKey).([]Goal); ok {
		return goals
	}
	return []Goal{}
}
