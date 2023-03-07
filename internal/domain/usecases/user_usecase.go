package usecases

import (
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

func (c *UserUseCase) CreateUser(input *dtos.UserInputDTO) (*dtos.UserOutputDTO, error) {
	// isExist, err := c.UserRepository.FindByEmail(input.Email)
	// if isExist.Email == input.Email {
	// 	return nil, errors.New("email already exists")
	// }

	// if err != nil {
	// 	return nil, err
	// }

	user, err := entities.NewUser(input.FullName, input.Email, input.ProfilePicture)
	if err != nil {
		return nil, err
	}

	if err := c.userRepository.CreateUser(user); err != nil {
		return nil, err
	}

	dto := dtos.UserOutputDTO{
		FullName:       user.FullName,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
	}

	return &dto, nil
}
