package vcago

import (
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Model string
}

func NewHandler(model string) *Handler {
	return &Handler{
		Model: model,
	}
}

func (i *Handler) Context(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := Context{
			i.Model,
			c,
		}
		return next(cc)
	}
}
