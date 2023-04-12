package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/reangeline/workout-plan-go/internal/domain/entities"
	"github.com/reangeline/workout-plan-go/internal/infra/database/postgres/sqlc"
)

type ExerciseRepository struct {
	dbConn *sql.DB
	sqlc   *sqlc.Queries
}

func NewExerciseRepository(dbConn *sql.DB) *ExerciseRepository {
	return &ExerciseRepository{
		dbConn: dbConn,
		sqlc:   sqlc.New(dbConn),
	}
}

func (u *ExerciseRepository) callTx(ctx context.Context, fn func(*sqlc.Queries) error) error {
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

func (u *ExerciseRepository) CreateExercise(ctx context.Context, exercise *entities.Exercise) error {

	err := u.callTx(ctx, func(q *sqlc.Queries) error {
		var err error

		err = q.CreateExercise(ctx, sqlc.CreateExerciseParams{
			IDExercise:   exercise.IDExercise,
			NameExercise: exercise.NameExercise,
			Description:  exercise.Description,
			UriGif:       exercise.UriGif,
		})

		if err != nil {
			return err
		}

		for _, value := range exercise.MuscleGroup {
			err = q.CreateMuscleGroupExercises(ctx, sqlc.CreateMuscleGroupExercisesParams{
				IDExercise:  exercise.IDExercise,
				MuscleGroup: value,
			})

			if err != nil {
				return err
			}
		}

		return nil

	})

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
