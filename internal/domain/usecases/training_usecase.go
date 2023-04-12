package usecases

import (
	"context"

	"github.com/reangeline/workout-plan-go/internal/domain/contracts/repositories"
	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type TrainingUseCase struct {
	exerciseRepository repositories.TrainingRepositoryInterface
}

func NewTrainingUseCase(exerciseRepository repositories.TrainingRepositoryInterface) *TrainingUseCase {
	return &TrainingUseCase{exerciseRepository}
}

func (u *TrainingUseCase) CreateTraining(ctx context.Context, input *dtos.CreateTrainingInput) error {

	exercise, err := entities.NewTraining(input.IDUser, input.TrainigName)
	if err != nil {
		return err
	}

	if err := u.exerciseRepository.CreateTraining(ctx, exercise); err != nil {
		return err
	}

	return nil
}
