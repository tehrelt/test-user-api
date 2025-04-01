package handlers

import (
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/tehrelt/test-users-api/internal/common"
	"github.com/tehrelt/test-users-api/internal/service"
	"github.com/tehrelt/test-users-api/internal/service/userservice"
)

func FindUser(userService *userservice.UserService) echo.HandlerFunc {

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

		rawId := c.Param("id")
		id, err := uuid.Parse(rawId)
		if err != nil {
			l.Warn("invalid id")
			return badrequest("invalid id")
		}

		ctx := common.InjectLogger(c.Request().Context(), l)

		user, err := userService.Find(ctx, id)
		if err != nil {
			if errors.Is(err, service.ErrUserNotFound) {
				l.Warn("user not found", slog.String("id", id.String()))
				return notfound("user not found")
			}

			l.Error("failed to find user", slog.String("err", err.Error()))
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
