package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sockstack/9c-cloud/auth/contract/service"
	"github.com/sockstack/9c-cloud/auth/model"
	service2 "github.com/sockstack/9c-cloud/auth/service"
	"net/http"
)

type UserHandler struct {
	service service.IUserService
	session service.ISessionService
}

func NewUserHandler() *UserHandler {
	return &UserHandler{service: service2.NewUserService(), session: service2.NewSessionService()}
}

func (u *UserHandler) LoginView(ctx *gin.Context) {
	session, _ := u.session.Get(ctx.Request, "rf")
	ctx.HTML(http.StatusOK, "login.html", session.Values["rf"])
}

func (u *UserHandler) DoLogin(ctx *gin.Context) {
	query := model.NewUserQuery()
	err := ctx.ShouldBind(query)
}
