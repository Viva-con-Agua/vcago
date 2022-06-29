package vcago

import (
	"github.com/labstack/echo/v4"
)

// Handler represents an network handler for echo framework
type Handler struct {
	Model string
}

//NewHandler creates an new `Handler`.
func NewHandler(model string) *Handler {
	return &Handler{
		Model: model,
	}
}

//Conext return an function for convertion an echo.Context models to an Context model.
//Based on the echo.HandlerFunc.
func (i *Handler) Context(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := Context{
			i.Model,
			c,
		}
		return next(cc)
	}
}
