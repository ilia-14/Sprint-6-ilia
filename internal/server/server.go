package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	*http.Server
}

// NewServer создает новую структуру приложения с настроенным сервером.
func NewServer(logger *log.Logger) *Server {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	app := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{Logger: logger, Server: app}
}
