// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: training_query.sql

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

const createTraining = `-- name: CreateTraining :exec
INSERT INTO trainings (id_training, id_user, trainig_name) 
VALUES ($1,$2,$3)
RETURNING id_training, id_user, trainig_name
`

type CreateTrainingParams struct {
	IDTraining  uuid.UUID
	IDUser      uuid.UUID
	TrainigName string
}

func (q *Queries) CreateTraining(ctx context.Context, arg CreateTrainingParams) error {
	_, err := q.db.ExecContext(ctx, createTraining, arg.IDTraining, arg.IDUser, arg.TrainigName)
	return err
}