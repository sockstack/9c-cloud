package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sockstack/9c-cloud/auth/handler"
	"net/http"
)

func authApi(e *gin.Engine)  {
	e.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"code": 200, "message": "ok"})
	})

	oauthHandler := handler.NewOauthHandler()
	e.POST("/token", oauthHandler.Token)
	e.GET("/authorize", oauthHandler.Authorize)
}
