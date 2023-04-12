package repositories

import (
	"context"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
)

type TrainingRepositoryInterface interface {
	CreateTraining(ctx context.Context, exercise *entities.Training) error
}
