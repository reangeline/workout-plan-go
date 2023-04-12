package usecases_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/reangeline/workout-plan-go/internal/domain/usecases"
	"github.com/reangeline/workout-plan-go/internal/dtos"
	database_test "github.com/reangeline/workout-plan-go/tests/unit/internal/infra/database"
)

func TestExerciseUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()

	repo := database_test.NewMemoryExerciseRepository()
	usecase := usecases.NewExerciseUseCase(repo)

	input := &dtos.CreateExerciseInput{
		NameExercise: "Supino Reto",
		MuscleGroup:  []string{"peito"},
		Description:  "Supino no banco com barra",
		UriGif:       "http://teste.gif",
	}

	t.Run("should create user successfully", func(t *testing.T) {
		err := usecase.CreateExercise(ctx, input)
		assert.NoError(t, err)
	})

}
