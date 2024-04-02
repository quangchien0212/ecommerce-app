package server

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/quangchien0212/ecommerce-app/internal/generic/common_errors"
	"github.com/quangchien0212/ecommerce-app/internal/models"
)

func (s *EchoServer) AddCategory(ctx echo.Context) error {
	category := new(models.Category)
	if err := ctx.Bind(category); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, map[string]any{"error": err.Error()})
	}
	category, err := s.DB.AddCategory(ctx.Request().Context(), category)
	if err != nil {
		var conflictError *common_errors.ConflictError
		var violationError *common_errors.ViolationError
		switch {
		case errors.As(err, &conflictError):
			return ctx.JSON(http.StatusConflict, map[string]any{"error": err.Error()})
		case errors.As(err, &violationError):
			return ctx.JSON(http.StatusConflict, map[string]any{"error": err.Error()})
		default:
			return ctx.JSON(http.StatusInternalServerError, map[string]any{"error": err.Error()})
		}
	}
	return ctx.JSON(http.StatusCreated, category)
}