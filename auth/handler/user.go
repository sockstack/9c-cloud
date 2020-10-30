package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sockstack/9c-cloud/auth/contract/service"
	service2 "github.com/sockstack/9c-cloud/auth/service"
	"net/http"
)

type UserHandler struct {
	service service.IUserService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{service: service2.NewUserService()}
}

func (u *UserHandler) LoginView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

