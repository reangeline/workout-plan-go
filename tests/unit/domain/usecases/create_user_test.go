package usecases_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/reangeline/workout-plan-go/internal/domain/entities"

	mocks_test "github.com/reangeline/workout-plan-go/tests/unit/domain/mocks"
)

func TestWhenCreateUser_ThenShouldReceiveAnNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newMockUser := mocks_test.NewMockUserRepositoryInterface(ctrl)

	newMockUser.EXPECT().Create(&entities.User{
		FullName:       "renato angeline",
		Email:          "reangeline@hotmail.com",
		ProfilePicture: "none",
	}).Return(nil).AnyTimes()

}

func TestWhenFindUserByEmail_ThenShouldReceiveAnEntityUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newMockUser := mocks_test.NewMockUserRepositoryInterface(ctrl)
	email := "reangeline@hotmail.com"
	newMockUser.EXPECT().FindByEmail(email).Return(entities.User{
		FullName:       "renato angeline",
		Email:          "reangeline@hotmail.com",
		ProfilePicture: "none",
	}, nil).AnyTimes()

}
