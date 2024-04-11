package handler

import (
	_ "github.com/SJ22032003/go-fullstack/model"
	util "github.com/SJ22032003/go-fullstack/util"
	favourites_views "github.com/SJ22032003/go-fullstack/views/favourites"
	gin "github.com/gin-gonic/gin"
	"net/http"
)

type FavouriteHandler struct{}

func (f *FavouriteHandler) Favourites(ctx *gin.Context) {
	template := util.NewTempl(ctx, http.StatusOK, favourites_views.FavouritePage())
	ctx.Render(http.StatusOK, template)
}
