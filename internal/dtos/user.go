package dtos

import "github.com/reangeline/workout-plan-go/pkg/entities"

type UserInputDTO struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	ProfilePicture string `json:"profile_picture"`
}

type UserOutputDTO struct {
	ID             entities.ID `json:"id"`
	FullName       string      `json:"full_name"`
	Email          string      `json:"email"`
	ProfilePicture string      `json:"profile_picture"`
}
