package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/infra/database/postgres/sqlc"
)

type TrainingRepository struct {
	dbConn *sql.DB
	sqlc   *sqlc.Queries
}

func NewTrainingRepository(dbConn *sql.DB) *TrainingRepository {
	return &TrainingRepository{
		dbConn: dbConn,
		sqlc:   sqlc.New(dbConn),
	}
}

func (u *TrainingRepository) callTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
	tx, err := u.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf("error on rollback: %v, original error: %w", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (u *TrainingRepository) CreateTraining(ctx context.Context, exercise *entities.Training) error {

	err := u.callTx(ctx, func(q *sqlc.Queries) error {

		err := q.CreateTraining(ctx, sqlc.CreateTrainingParams{
			IDTraining: exercise.TrainingID,
			IDUser:     exercise.UserID,
		})

		if err != nil {
			return err
		}

		return nil

	})

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
