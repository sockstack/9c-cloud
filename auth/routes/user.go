package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sockstack/9c-cloud/auth/handler"
)

func userApi(e *gin.Engine)  {
	userHandler := handler.NewUserHandler()
	e.GET("/login", userHandler.LoginView)
}
