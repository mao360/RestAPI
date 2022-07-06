package repository

import "github.com/jackc/pgx/v4"

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB() (*pgx., error) { // возвращает бд и ошибку
	db, err := pgx.Begin()
}
