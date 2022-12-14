// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: Driver.sql

package pgsql

import (
	"context"
	"database/sql"
)

const assignEngine = `-- name: AssignEngine :exec
UPDATE assets.drivers
SET unit_id =$2
WHERE id = $1
`

type AssignEngineParams struct {
	ID     int64         `json:"id"`
	UnitID sql.NullInt32 `json:"unit_id"`
}

func (q *Queries) AssignEngine(ctx context.Context, arg AssignEngineParams) error {
	_, err := q.db.ExecContext(ctx, assignEngine, arg.ID, arg.UnitID)
	return err
}

const createEngineDriver = `-- name: CreateEngineDriver :exec
INSERT INTO assets.drivers
(engine_id, unit_id)
VALUES
($1, $2)
`

type CreateEngineDriverParams struct {
	EngineID sql.NullInt32 `json:"engine_id"`
	UnitID   sql.NullInt32 `json:"unit_id"`
}

func (q *Queries) CreateEngineDriver(ctx context.Context, arg CreateEngineDriverParams) error {
	_, err := q.db.ExecContext(ctx, createEngineDriver, arg.EngineID, arg.UnitID)
	return err
}

const createMotorDriver = `-- name: CreateMotorDriver :exec
INSERT INTO assets.drivers
(motor_id, unit_id)
VALUES
($1, $2)
`

type CreateMotorDriverParams struct {
	MotorID sql.NullInt32 `json:"motor_id"`
	UnitID  sql.NullInt32 `json:"unit_id"`
}

func (q *Queries) CreateMotorDriver(ctx context.Context, arg CreateMotorDriverParams) error {
	_, err := q.db.ExecContext(ctx, createMotorDriver, arg.MotorID, arg.UnitID)
	return err
}

const deletDriver = `-- name: DeletDriver :exec
DELETE FROM assets.drivers
WHERE id = $1
`

func (q *Queries) DeletDriver(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletDriver, id)
	return err
}

const getDriverById = `-- name: GetDriverById :one
SELECT id, engine_id, motor_id, unit_id, created_at, modified_at FROM assets.Drivers
WHERE id = $1
`

func (q *Queries) GetDriverById(ctx context.Context, id int64) (AssetsDriver, error) {
	row := q.db.QueryRowContext(ctx, getDriverById, id)
	var i AssetsDriver
	err := row.Scan(
		&i.ID,
		&i.EngineID,
		&i.MotorID,
		&i.UnitID,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getDriverByUnitID = `-- name: GetDriverByUnitID :one
SELECT id, engine_id, motor_id, unit_id, created_at, modified_at FROM assets.drivers
WHERE unit_id = $1
`

func (q *Queries) GetDriverByUnitID(ctx context.Context, unitID sql.NullInt32) (AssetsDriver, error) {
	row := q.db.QueryRowContext(ctx, getDriverByUnitID, unitID)
	var i AssetsDriver
	err := row.Scan(
		&i.ID,
		&i.EngineID,
		&i.MotorID,
		&i.UnitID,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
