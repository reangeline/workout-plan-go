package usecases

import (
	"context"

	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type ExerciseUseCaseInterface interface {
	CreateExercise(ctx context.Context, input *dtos.CreateExerciseInput) error
}
