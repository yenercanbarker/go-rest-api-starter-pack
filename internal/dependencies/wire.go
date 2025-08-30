//go:build wireinject
// +build wireinject

package dependencies

import (
	"github.com/google/wire"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/config"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/handlers"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/repositories"
	"github.com/yenercanbarker/go-rest-api-starter-pack/internal/services"
)

var CommonSet = wire.NewSet(
	config.Load,
	config.InitDatabase,
)

func InitUserDependencyInjection() (*handlers.UserHandler, error) {
	wire.Build(
		CommonSet,
		wire.Bind(new(repositories.UserRepositoryInterface), new(*repositories.UserRepository)),
		wire.Bind(new(services.UserServiceInterface), new(*services.UserService)),
		repositories.NewUserRepository,
		services.NewUserService,
		handlers.NewUserHandler,
	)
	return &handlers.UserHandler{}, nil
}
