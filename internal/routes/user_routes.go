package routes

import (
	"SpendSmartAPI/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	users := r.Group("/users")
	{
		users.POST("/", userHandler.Create)
		users.GET("/", userHandler.FindAll)
		users.GET("/:id", userHandler.FindById)
	}
}
