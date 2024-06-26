package handler

import (
	"fmt"
	"net/http"
	"strings"

	util "github.com/SJ22032003/go-fullstack/util"
	login_view "github.com/SJ22032003/go-fullstack/views/authentication/login"
	gin "github.com/gin-gonic/gin"
)

type AuthHandler struct{}

func (a *AuthHandler) AuthPage(ctx *gin.Context) {
	template := util.NewTempl(ctx, http.StatusOK, login_view.LoginPage())
	ctx.Render(http.StatusOK, template)
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	var errors []string

	if email == "" {
		errors = append(errors, "Email is required")
	}

	if password == "" || len(password) < 8 {
		errors = append(errors, "Password is required")
	}

	if len(errors) > 0 {
		template := util.NewTempl(ctx, http.StatusOK, login_view.AuthErrorCmp(errors))
		ctx.Render(http.StatusOK, template)
		return
	}
	ctx.Header("HX-Redirect", "/goals")
	ctx.String(http.StatusFound, "")

}

func (a *AuthHandler) InputValidation(ctx *gin.Context) {

	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	fmt.Println(email, password, "here")

	if email != "" && !strings.Contains(email, "@") {
		ctx.String(http.StatusOK, "Email is required")
		return
	}

	if password != "" && len(password) < 8 {
		ctx.String(http.StatusOK, "Password is required and must be at least 8 characters long")
		return
	}
	ctx.String(http.StatusOK, "")

}
