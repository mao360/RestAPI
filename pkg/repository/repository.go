package repository

import "github.com/jackc/pgx/v4/pgxpool"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{}
}
