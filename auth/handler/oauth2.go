package handler

import (
	"github.com/gin-gonic/gin"
	service2 "github.com/sockstack/9c-cloud/auth/contract/service"
	"github.com/sockstack/9c-cloud/auth/service"
)

type OauthHandler struct {
	oauth service2.IOauth2
}

func NewOauthHandler() *OauthHandler {
	return &OauthHandler{oauth: service.NewOauth2()}
}

func (o *OauthHandler) Authorize(ctx *gin.Context)  {
	o.oauth.Unwrap(o.oauth.HandleTokenRequest(ctx.Writer, ctx.Request))
}

func (o *OauthHandler) Token(ctx *gin.Context)  {
	o.oauth.Unwrap(o.oauth.HandleTokenRequest(ctx.Writer, ctx.Request))
}
