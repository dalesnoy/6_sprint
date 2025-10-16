package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	logger := log.New(os.Stdout, "[SERVER]", log.LstdFlags)
	srv := server.New(logger)

	logger.Println("Сервер запускается на http://localhost:8080")

	if err := srv.HTTPserver.ListenAndServe(); err != nil && err != server.ErrServerClosed {
		logger.Fatal("Ошибка сервера: ", err)
	}
}
