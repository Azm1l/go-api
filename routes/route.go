package routes

import (
	"github.com/Azm1l/rest-api-go/handler"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userHandler *handler.UserHandler) {
	user := r.Group("/user")
	{
		user.POST("/", userHandler.CreateUser)
		user.GET("/", userHandler.ShowAllUsers)
		user.GET("/:id", userHandler.FindUserByID)
		user.PUT("/:id", userHandler.UpdateUser)
	}
}
