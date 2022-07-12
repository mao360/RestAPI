package main

import (
	"context"
	"github.com/mao360/RestAPI"
	"github.com/mao360/RestAPI/pkg/handler"
	"github.com/mao360/RestAPI/pkg/repository"
	"github.com/mao360/RestAPI/pkg/service"
	"log"
)

//ЗАПУСК СЕРВЕРА
func main() {
	ctx, cancel := context.Background() // крч тут надо фиксить
	defer cancel
	db, err := repository.NewClient(ctx, repository.Config{
		Host:     "localhost",
		Port:     "5436",
		UserName: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	}) // зачем контекст непонятно
	if err != nil {
		log.Fatal("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(RestAPI.Server)
	err := srv.Run("8000", handlers.InitRouts())
	if err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
