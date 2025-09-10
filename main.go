package main

import (
	"log"

	"github.com/Azm1l/rest-api-go/config"
	"github.com/Azm1l/rest-api-go/entity"
	"github.com/Azm1l/rest-api-go/handler"
	"github.com/Azm1l/rest-api-go/repository"
	"github.com/Azm1l/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	// auto migrate
	db.AutoMigrate(&entity.User{})

	//init layer
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	user := r.Group("user")
	user.POST("/", userHandler.CreateUser)
	user.GET("/", userHandler.ShowAllUsers)
	user.GET("/:id", userHandler.FindUserByID)
	r.Run(":8080")
}
