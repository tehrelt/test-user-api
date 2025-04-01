package middlewares

import (
	"log/slog"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tehrelt/test-users-api/internal/common"
)

func RequestIdMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		requestId := uuid.New().String()
		l := slog.With(slog.String("requestId", requestId))

		return func(c echo.Context) error {
			c.Set(string(common.LogKey), l)
			c.Response().Header().Set("X-Request-Id", requestId)

			return next(c)
		}
	}
}
