package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rchenhyy/demo-ginex/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	root := router.Group("/")
	{
		root.Any("/", handlers.HelloGin)
	}

	user := router.Group("/user/")
	{
		user.GET("/:name", handlers.GetUser)
		user.POST("/register", handlers.RegisterUser)
	}
	return router
}
