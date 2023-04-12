package usecases

import (
	"context"

	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type TrainingUseCaseInterface interface {
	CreateTraining(ctx context.Context, input *dtos.CreateTrainingInput) error
}
