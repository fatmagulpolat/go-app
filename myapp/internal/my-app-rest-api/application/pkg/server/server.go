package server

import (
	"context"
	"fmt"
	"myapp/configuration"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

type server struct {
	app *fiber.App
}

func NewServer(app *fiber.App) *server {
	return &server{
		app: app,
	}
}

// metot isminden önce yazılan şey hangi structa metot olarak yazıldığnı belirtiyor func (s *server)
func (s *server) StartHttpServer() {
	go func() {
		gracefulShutdown(s.app)
	}()
	if err := s.app.Listen(fmt.Sprintf(":%s", configuration.Port)); err != nil && err != http.ErrServerClosed {
		fmt.Printf("cannot start server - ERROR : %v\n", err)
		panic("cannot start server")
	}
}

func gracefulShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown server")

	_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		fmt.Printf("Server shutdown")
	}

	fmt.Println("Server exiting")
}
