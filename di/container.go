package di

import (
	"todo-app-go/application/user"
	"todo-app-go/infrastructure/authclient"
	"todo-app-go/infrastructure/database"
	"todo-app-go/infrastructure/database/repository"
	"todo-app-go/interface/handler"

	"github.com/google/wire"
)

// UserControllerの依存関係組み立て
func InitializedUserController() (*handler.UserHandler, error) {
	wire.Build(
		database.NewGormDB,
		repository.NewUserRepositoryProvider,
		handler.NewUserHandler,
		user.NewLoginUsecase,
		user.NewRegisterUsecase,
		authclient.NewAuthClient,
	)
	return &handler.UserHandler{}, nil
}
