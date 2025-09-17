package main

import (
	"log"
	"time"

	"github.com/Azm1l/rest-api-go/bootstrap"
	"github.com/Azm1l/rest-api-go/config"
	"github.com/Azm1l/rest-api-go/entity"
	"github.com/Azm1l/rest-api-go/middleware"
	"github.com/Azm1l/rest-api-go/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := config.ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	// auto migrate
	db.AutoMigrate(&entity.User{})

	//dependencies
	deps := bootstrap.InitDependencies(db)

	r := gin.Default()

	//middleware
	r.Use(middleware.SlowRequestLogger(500 * time.Millisecond))

	routes.UserRoutes(r, deps.UserHandler)
	r.Run(":8080")
}
