package http

import (
	"context"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/tehrelt/test-users-api/internal/config"
	"github.com/tehrelt/test-users-api/internal/service/userservice"
	"github.com/tehrelt/test-users-api/internal/transport/http/handlers"
	"github.com/tehrelt/test-users-api/internal/transport/http/middlewares"
)

type Server struct {
	cfg         *config.Config
	router      *echo.Echo
	userService *userservice.UserService
}

func New(cfg *config.Config, userService *userservice.UserService) *Server {

	s := &Server{
		cfg:         cfg,
		router:      echo.New(),
		userService: userService,
	}

	s.setup()

	return s
}

func (s *Server) setup() {

	s.router.Use(middlewares.RequestIdMiddleware())
	s.router.Use(middlewares.LoggingMiddleware())

	userGroup := s.router.Group("/users")
	userGroup.POST("/", handlers.CreateUser(s.userService))
	userGroup.GET("/:id", handlers.FindUser(s.userService))
	userGroup.PUT("/:id", handlers.UpdateUser(s.userService))
}

func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.cfg.Http.Host, s.cfg.Http.Port)
	return s.router.Start(addr)
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.router.Shutdown(ctx)
}
