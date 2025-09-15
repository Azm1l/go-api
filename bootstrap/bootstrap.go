package bootstrap

import (
	"github.com/Azm1l/rest-api-go/handler"
	"github.com/Azm1l/rest-api-go/repository"
	"github.com/Azm1l/rest-api-go/service"
	"gorm.io/gorm"
)

type AppDependencies struct {
	UserHandler *handler.UserHandler
	// kalau nanti ada GameHandler, TransactionHandler, tambahin di sini juga
}

func InitDependencies(db *gorm.DB) *AppDependencies {
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	return &AppDependencies{
		UserHandler: userHandler,
	}
}
