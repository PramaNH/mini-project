// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES ($1, $2)
RETURNING id, username
`

type CreateUserParams struct {
	Username string
	Password string
}

type CreateUserRow struct {
	ID       int32
	Username string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password)
	var i CreateUserRow
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, password
FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}
