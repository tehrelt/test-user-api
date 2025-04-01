package middlewares

import (
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tehrelt/test-users-api/internal/common"
)

func LoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			start := time.Now()
			l := c.Get(string(common.LogKey)).(*slog.Logger)

			defer func() {
				l = l.With(
					slog.String("method", c.Request().Method),
					slog.String("path", c.Path()),
					slog.Int("statusCode", c.Response().Status),
					slog.Duration("latency", time.Since(start)),
				)

				if err != nil {
					l.Error("request error", slog.String("err", err.Error()))
					return
				}

				l.Info("request completed")
			}()

			err = next(c)
			return err
		}
	}
}
