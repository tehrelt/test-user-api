package handlers

import (
	"errors"
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tehrelt/test-users-api/internal/common"
	"github.com/tehrelt/test-users-api/internal/service"
	"github.com/tehrelt/test-users-api/internal/service/userservice"
)

func CreateUser(userService *userservice.UserService) echo.HandlerFunc {

	type request struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}

	type response struct {
		Id        string     `json:"id"`
		FirstName string     `json:"firstName"`
		LastName  string     `json:"lastName"`
		Email     string     `json:"email"`
		CreatedAt time.Time  `json:"createdAt"`
		UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	}

	return func(c echo.Context) error {
		l := extractLogger(c)

		var req request
		if err := c.Bind(&req); err != nil {
			l.Error("failed to bind request", slog.String("err", err.Error()))
			return badrequest("failed validation of request")
		}

		ctx := common.InjectLogger(c.Request().Context(), l)

		user, err := userService.Create(ctx, &service.CreateUserDto{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
		})
		if err != nil {
			if errors.Is(err, service.ErrUserAlreadyExists) {
				l.Warn("user already exists")
				return badrequest("user already exists")
			}

			l.Error("failed to create user", slog.String("err", err.Error()))
			return internal("internal server error")
		}

		return ok(c, &response{
			Id:        user.Id.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
}
