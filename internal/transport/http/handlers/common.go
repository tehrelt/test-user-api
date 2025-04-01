package handlers

import (
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/tehrelt/test-users-api/internal/common"
)

func extractLogger(e echo.Context) *slog.Logger {
	return e.Get(string(common.LogKey)).(*slog.Logger)
}

func notfound(msg string) error {
	return echo.NewHTTPError(echo.ErrNotFound.Code, ErrorMessage{
		Code:    echo.ErrNotFound.Code,
		Message: msg,
	})
}

func badrequest(msg string) error {
	return echo.NewHTTPError(echo.ErrBadRequest.Code, ErrorMessage{
		Code:    echo.ErrBadRequest.Code,
		Message: msg,
	})
}

func internal(msg string) error {
	return echo.NewHTTPError(echo.ErrInternalServerError.Code, ErrorMessage{
		Code:    echo.ErrInternalServerError.Code,
		Message: msg,
	})
}

func ok(c echo.Context, val interface{}) error {
	return c.JSON(200, val)
}
