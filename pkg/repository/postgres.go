package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Config struct { // duno for what this structure
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func NewClient(ctx context.Context, cfg Config) (pool *pgxpool.Pool, err error) { //возможно не работает правильно, не понятно зачем тут контест в арг
	var maxAttempts = 4
	dsn := fmt.Sprintf("postgresql://host:%s port:%s username:%s password:%s sslmode:%s", cfg.Host, cfg.Port, cfg.UserName, cfg.Password, cfg.SSLMode)
	for maxAttempts > 0 {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		pool, err = pgxpool.Connect(ctx, dsn)
		if err == nil {
			return pool, nil
		}
		maxAttempts--
	}
	if err != nil {
		fmt.Print("failed to connect to postgresql")
		return nil, err
	}
	return pool, nil
}
