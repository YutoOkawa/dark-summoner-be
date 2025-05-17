package server

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	App             *fiber.App
	port            string
	shutdownTimeout time.Duration
}

func NewServer(port string, shutdownTimeout time.Duration) *Server {
	app := fiber.New()
	return &Server{
		App:             app,
		port:            port,
		shutdownTimeout: shutdownTimeout,
	}
}

func (s *Server) Start(errCh chan<- error) {
	if err := s.App.Listen(s.port); err != nil {
		errCh <- err
	}
}

func (s *Server) Shutdown() error {
	return s.App.ShutdownWithTimeout(s.shutdownTimeout)
}
