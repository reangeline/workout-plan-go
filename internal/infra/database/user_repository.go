package database

import (
	"database/sql"
	"fmt"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/dtos"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (u *UserRepository) CreateUser(user *entities.User) error {
	fmt.Println(user)
	return nil
}

func (u *UserRepository) FindByEmail(email string) (dtos.UserOutputDTO, error) {
	return dtos.UserOutputDTO{}, nil
}
