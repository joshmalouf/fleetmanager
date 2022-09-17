// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: Motor.sql

package pgsql

import (
	"context"
	"database/sql"
)

const assignMotor = `-- name: AssignMotor :exec
UPDATE assets.Motors
SET unit_id = $2
WHERE id = $1
`

type AssignMotorParams struct {
	ID     int64         `json:"id"`
	UnitID sql.NullInt32 `json:"unit_id"`
}

func (q *Queries) AssignMotor(ctx context.Context, arg AssignMotorParams) error {
	_, err := q.db.ExecContext(ctx, assignMotor, arg.ID, arg.UnitID)
	return err
}

const deactivateMotor = `-- name: DeactivateMotor :exec
UPDATE assets.Motors
SET opstatus = "inactive", unit_id = NULL
WHERE id = $1
`

func (q *Queries) DeactivateMotor(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deactivateMotor, id)
	return err
}

const getAvailMotors = `-- name: GetAvailMotors :many
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Motors
WHERE opstatus = "active" AND unit_id = NULL
`

func (q *Queries) GetAvailMotors(ctx context.Context) ([]AssetsMotor, error) {
	rows, err := q.db.QueryContext(ctx, getAvailMotors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AssetsMotor{}
	for rows.Next() {
		var i AssetsMotor
		if err := rows.Scan(
			&i.ID,
			&i.Make,
			&i.Model,
			&i.SerialNumber,
			&i.UnitID,
			&i.OpStatus,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMotorByID = `-- name: GetMotorByID :one
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Motors
WHERE id = $1
`

func (q *Queries) GetMotorByID(ctx context.Context, id int64) (AssetsMotor, error) {
	row := q.db.QueryRowContext(ctx, getMotorByID, id)
	var i AssetsMotor
	err := row.Scan(
		&i.ID,
		&i.Make,
		&i.Model,
		&i.SerialNumber,
		&i.UnitID,
		&i.OpStatus,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getMotorByMakeModel = `-- name: GetMotorByMakeModel :many
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Motors
WHERE make = $1 AND model = $2
ORDER BY throws
`

type GetMotorByMakeModelParams struct {
	Make  string `json:"make"`
	Model string `json:"model"`
}

func (q *Queries) GetMotorByMakeModel(ctx context.Context, arg GetMotorByMakeModelParams) ([]AssetsMotor, error) {
	rows, err := q.db.QueryContext(ctx, getMotorByMakeModel, arg.Make, arg.Model)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AssetsMotor{}
	for rows.Next() {
		var i AssetsMotor
		if err := rows.Scan(
			&i.ID,
			&i.Make,
			&i.Model,
			&i.SerialNumber,
			&i.UnitID,
			&i.OpStatus,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMotorBySerial = `-- name: GetMotorBySerial :one
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Motors
WHERE serial_number = $1
`

func (q *Queries) GetMotorBySerial(ctx context.Context, serialNumber string) (AssetsMotor, error) {
	row := q.db.QueryRowContext(ctx, getMotorBySerial, serialNumber)
	var i AssetsMotor
	err := row.Scan(
		&i.ID,
		&i.Make,
		&i.Model,
		&i.SerialNumber,
		&i.UnitID,
		&i.OpStatus,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getMotors = `-- name: GetMotors :many
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Motors
ORDER BY make
`

func (q *Queries) GetMotors(ctx context.Context) ([]AssetsMotor, error) {
	rows, err := q.db.QueryContext(ctx, getMotors)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AssetsMotor{}
	for rows.Next() {
		var i AssetsMotor
		if err := rows.Scan(
			&i.ID,
			&i.Make,
			&i.Model,
			&i.SerialNumber,
			&i.UnitID,
			&i.OpStatus,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMotorsByMake = `-- name: GetMotorsByMake :many
SELECT id, make, model, serial_number, unit_id, op_status, created_at, modified_at FROM assets.Motors
WHERE make = $1
ORDER BY model
`

func (q *Queries) GetMotorsByMake(ctx context.Context, make string) ([]AssetsMotor, error) {
	rows, err := q.db.QueryContext(ctx, getMotorsByMake, make)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AssetsMotor{}
	for rows.Next() {
		var i AssetsMotor
		if err := rows.Scan(
			&i.ID,
			&i.Make,
			&i.Model,
			&i.SerialNumber,
			&i.UnitID,
			&i.OpStatus,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}