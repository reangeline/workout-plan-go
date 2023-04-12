package usecases

import (
	"context"

	"github.com/reangeline/workout-plan-go/internal/domain/contracts/repositories"
	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type ExerciseUseCase struct {
	exerciseRepository repositories.ExerciseRepositoryInterface
}

func NewExerciseUseCase(
	exerciseRepository repositories.ExerciseRepositoryInterface,
) *ExerciseUseCase {
	return &ExerciseUseCase{exerciseRepository}
}

func (u *ExerciseUseCase) CreateExercise(ctx context.Context, input *dtos.CreateExerciseInput) error {
	exercise, err := entities.NewExercise(input.NameExercise, input.MuscleGroup, input.Description, input.UriGif)
	if err != nil {
		return err
	}

	if err := u.exerciseRepository.CreateExercise(ctx, exercise); err != nil {
		return err
	}

	return nil
}
