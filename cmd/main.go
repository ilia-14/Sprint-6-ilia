package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаем логгер
	logger := log.New(os.Stdout, "myapp:", log.LstdFlags)

	app := server.NewServer(logger)

	// Логируем старт сервера
	logger.Printf("Сервер запущен на адресе: %s\n", app.Addr)

	err := app.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
