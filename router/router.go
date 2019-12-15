package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rchenhyy/demo-ginex/handlers"
	"github.com/rchenhyy/demo-ginex/middlewares"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.Logger())

	user := router.Group("/user/")
	{
		user.GET("/:name", handlers.User)
		user.POST("/register", handlers.UserRegister)
		user.POST("/login", handlers.UserLogin)
	}

	root := router.Group("/", middlewares.CheckLogin()) // use middleware in specific paths
	{
		root.Any("/", handlers.HelloGin)
	}
	return router
}
