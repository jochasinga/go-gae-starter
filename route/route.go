package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jochasinga/boo/handler"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	base := router.Group("/v0")
	base.GET("/hello/:name", handler.HelloHandler)
	base.POST("/hello", handler.HelloPostHandler)

	{
		user := base.Group("/users")
		user.GET("", handler.GetUsersHandler)
		user.POST("", handler.CreateUserHandler)
		user.GET("/:id", handler.GetUserHandler)
	}
	return router
}
