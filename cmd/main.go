package main

import (
	"github.com/mao360/RestAPI"
	"github.com/mao360/RestAPI/pkg/handler"
	"github.com/mao360/RestAPI/pkg/repository"
	"github.com/mao360/RestAPI/pkg/service"
	"log"
)

//ЗАПУСК СЕРВЕРА
func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(RestAPI.Server)
	err := srv.Run("8000", handlers.InitRouts())
	if err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
