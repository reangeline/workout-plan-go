package usecases

import (
	"context"
	"errors"

	"github.com/reangeline/workout-plan-go/internal/domain/contracts/repositories"
	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type UserUseCase struct {
	userRepository repositories.UserRepositoryInterface
}

func NewUserUseCase(userRepository repositories.UserRepositoryInterface) *UserUseCase {
	return &UserUseCase{userRepository}
}

var (
	ErrEmailAlreadyExists = errors.New("email already exist")
)

func (u *UserUseCase) CreateUser(ctx context.Context, input *dtos.CreateUserInput) error {

	isExist, _ := u.CheckEmailExists(input.Email)

	if isExist {
		return ErrEmailAlreadyExists
	}

	user, err := entities.NewUser(input.FullName, input.Email, input.ProfilePicture)
	if err != nil {
		return err
	}

	if err := u.userRepository.CreateUser(ctx, user); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) CheckEmailExists(email string) (bool, error) {
	output, err := u.userRepository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if output.Email != email {
		return false, nil
	}

	return true, nil
}

func (u *UserUseCase) FindByEmail(email string) (*dtos.UserOutputDTO, error) {
	user, err := u.userRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil

}
