package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/dependencies"
)

func InitUserRoutes(router *gin.RouterGroup) {
	userHandler, err := dependencies.InitUserDependencyInjection()
	if err != nil {
		panic(err)
	}

	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}
}
