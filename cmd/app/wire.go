//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"database/sql"

	"github.com/reangeline/workout-plan-go/internal/domain/contracts/repositories"
	uc_interface "github.com/reangeline/workout-plan-go/internal/domain/contracts/usecases"
	"github.com/reangeline/workout-plan-go/internal/validation/protocols"

	"github.com/reangeline/workout-plan-go/internal/domain/usecases"
	"github.com/reangeline/workout-plan-go/internal/infra/database"
	"github.com/reangeline/workout-plan-go/internal/validation/validators"

	"github.com/reangeline/workout-plan-go/internal/infra/http/controllers"
)

var setUserRepositoryDependency = wire.NewSet(
	database.NewUserRepository,
	wire.Bind(new(repositories.UserRepositoryInterface), new(*database.UserRepository)),
)

var setUserUseCaseDependency = wire.NewSet(
	usecases.NewUserUseCase,
	wire.Bind(new(uc_interface.UserUseCaseInterface), new(*usecases.UserUseCase)),
)

var setUserValidatorDependency = wire.NewSet(
	validators.NewUserValidator,
	wire.Bind(new(protocols.UserValidatorInterface), new(*validators.UserValidator)),
)

func InitializeUserController(db *sql.DB) (*controllers.UserController, error) {
	wire.Build(
		setUserRepositoryDependency,
		setUserUseCaseDependency,
		setUserValidatorDependency,
		controllers.NewUserController,
	)
	return &controllers.UserController{}, nil
}
