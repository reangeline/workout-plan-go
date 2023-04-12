package repositories

import (
	"context"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
)

type ExerciseRepositoryInterface interface {
	CreateExercise(ctx context.Context, exercise *entities.Exercise) error
}
