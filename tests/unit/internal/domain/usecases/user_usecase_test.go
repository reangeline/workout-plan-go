package usecases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/reangeline/workout-plan-go/internal/domain/usecases"
	"github.com/reangeline/workout-plan-go/internal/dtos"
	database_test "github.com/reangeline/workout-plan-go/tests/unit/internal/infra/database"
)

func TestUserUseCase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := database_test.NewMemoryUserRepository()
	usecase := usecases.NewUserUseCase(repo)

	input := &dtos.UserInputDTO{
		FullName:       "Renato Angeline",
		Email:          "reangeline@hotmail.com",
		ProfilePicture: "teste",
	}

	t.Run("should create user successfully", func(t *testing.T) {
		err := usecase.CreateUser(input)
		assert.NoError(t, err)
	})

	t.Run("should find user by email", func(t *testing.T) {
		user, err := usecase.FindByEmail(input.Email)
		assert.NoError(t, err)
		assert.Equal(t, user.Email, input.Email)
	})

	t.Run("should not add a email that alrealdy exist", func(t *testing.T) {
		err := usecase.CreateUser(input)
		assert.Error(t, err)
	})

}
