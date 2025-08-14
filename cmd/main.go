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

	err := app.ListenAndServe()
	if err != nil {
		logger.Fatal(err)
	}
}
