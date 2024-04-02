package server

import (
	"log"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/quangchien0212/ecommerce-app/internal/database"
	"github.com/quangchien0212/ecommerce-app/internal/server/abstract"
)

type Server interface {
	Start() error
	abstract.Category
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DBClient
}

func NewServer(db database.DBClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (s *EchoServer) Start() error {
	slog.Info("serving at port 8080")
	err := s.echo.Start(":8080")
	if err != nil {
		log.Fatalf("Server Issue: %s", err)
		return err
	}
	return nil
}