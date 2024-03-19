// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: groupQuery.sql

package crud

import (
	"context"
	"time"
)

const checkIfCreator = `-- name: CheckIfCreator :one
SELECT COUNT(*) FROM group_
WHERE creator = ? AND id = ? LIMIT 1
`

type CheckIfCreatorParams struct {
	Creator int64
	ID      int64
}

func (q *Queries) CheckIfCreator(ctx context.Context, arg CheckIfCreatorParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, checkIfCreator, arg.Creator, arg.ID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createGroup = `-- name: CreateGroup :one
INSERT INTO group_ (
  title, creator, description_, created_at
) VALUES (
  ?, ?, ?, ?
)
RETURNING id, title, creator, description_, created_at
`

type CreateGroupParams struct {
	Title       string
	Creator     int64
	Description string
	CreatedAt   time.Time
}

func (q *Queries) CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error) {
	row := q.db.QueryRowContext(ctx, createGroup,
		arg.Title,
		arg.Creator,
		arg.Description,
		arg.CreatedAt,
	)
	var i Group
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Creator,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deleteGroup = `-- name: DeleteGroup :exec
DELETE FROM group_
WHERE id = ?
`

func (q *Queries) DeleteGroup(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteGroup, id)
	return err
}

const getAllGroups = `-- name: GetAllGroups :many
SELECT id, title, creator, description_, created_at FROM group_
`

func (q *Queries) GetAllGroups(ctx context.Context) ([]Group, error) {
	rows, err := q.db.QueryContext(ctx, getAllGroups)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Group
	for rows.Next() {
		var i Group
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Creator,
			&i.Description,
			&i.CreatedAt,
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

const getGroup = `-- name: GetGroup :one
SELECT id, title, creator, description_, created_at FROM group_
WHERE id = ? LIMIT 1
`

func (q *Queries) GetGroup(ctx context.Context, id int64) (Group, error) {
	row := q.db.QueryRowContext(ctx, getGroup, id)
	var i Group
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Creator,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}