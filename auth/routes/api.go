package routes

import "github.com/gin-gonic/gin"

func Api(engine *gin.Engine)  {
	authApi(engine)
}
