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

var setExerciseRepositoryDependency = wire.NewSet(
	database.NewExerciseRepository,
	wire.Bind(new(repositories.ExerciseRepositoryInterface), new(*database.ExerciseRepository)),
)

var setExerciseUseCaseDependency = wire.NewSet(
	usecases.NewExerciseUseCase,
	wire.Bind(new(uc_interface.ExerciseUseCaseInterface), new(*usecases.ExerciseUseCase)),
)

var setExerciseValidatorDependency = wire.NewSet(
	validators.NewExerciseValidator,
	wire.Bind(new(protocols.ExerciseValidatorInterface), new(*validators.ExerciseValidator)),
)

var setTrainingRepositoryDependency = wire.NewSet(
	database.NewTrainingRepository,
	wire.Bind(new(repositories.TrainingRepositoryInterface), new(*database.TrainingRepository)),
)

var setTrainingUseCaseDependency = wire.NewSet(
	usecases.NewTrainingUseCase,
	wire.Bind(new(uc_interface.TrainingUseCaseInterface), new(*usecases.TrainingUseCase)),
)

var setTrainingValidatorDependency = wire.NewSet(
	validators.NewTrainingValidator,
	wire.Bind(new(protocols.TrainingValidatorInterface), new(*validators.TrainingValidator)),
)

func InitializeExerciseController(db *sql.DB) (*controllers.ExerciseController, error) {
	wire.Build(
		setExerciseRepositoryDependency,
		setExerciseUseCaseDependency,
		setExerciseValidatorDependency,
		controllers.NewExerciseController,
	)
	return &controllers.ExerciseController{}, nil
}

func InitializeUserController(db *sql.DB) (*controllers.UserController, error) {
	wire.Build(
		setUserRepositoryDependency,
		setUserUseCaseDependency,
		setUserValidatorDependency,
		controllers.NewUserController,
	)
	return &controllers.UserController{}, nil
}

func InitializeTrainingController(db *sql.DB) (*controllers.TrainingController, error) {
	wire.Build(
		setTrainingRepositoryDependency,
		setTrainingUseCaseDependency,
		setTrainingValidatorDependency,
		controllers.NewTrainingController,
	)
	return &controllers.TrainingController{}, nil
}
