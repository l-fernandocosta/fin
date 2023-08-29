// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users( email, hashed_password ) VALUES ($1, $2) RETURNING id, email, hashed_password, created_at, updated_at
`

type CreateUserParams struct {
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, hashed_password, created_at, updated_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, hashed_password, created_at, updated_at FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
