package main

import (
	"github.com/mao360/RestAPI"
	"github.com/mao360/RestAPI/pkg/handler"
	"log"
)

//ЗАПУСК СЕРВЕРА
func main() {
	handlers := new(handler.Handler)
	srv := new(RestAPI.Server)
	err := srv.Run("8000", handlers.InitRouts())
	if err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
