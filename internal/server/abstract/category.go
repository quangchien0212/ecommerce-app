package abstract

import (
	"github.com/labstack/echo/v4"
)

type Category interface {
	AddCategory(ctx echo.Context) error
}