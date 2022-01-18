package verr

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseError struct {
	Message string `json:"message"`
}

const (
	//ColorRed use string(ColorRed)
	ColorRed = "\033[31m"
	//ColorGreen use string(ColorGreen)
	ColorGreen = "\033[32m"
	//ColorYellow use string(ColorYellow)
	ColorYellow = "\033[33m"
	//ColorBlue use string(ColorBlue)
	ColorBlue = "\033[34m"
	//ColorPurple use string(ColorPurple)
	ColorPurple = "\033[35m"
	//ColorCyan use string(ColorCyan)
	ColorCyan = "\033[36m"
	//ColorWhite use string(ColorWhite)
	ColorWhite = "\033[37m"
	//ErrorWithColor is a red string "Error: "
	ErrorWithColor = string(ColorRed) + "Error: " + string(ColorWhite)
	//WarningWithColor is a yellow string "Warning: "
	WarningWithColor = string(ColorYellow) + "Warning: " + string(ColorWhite)
	//SuccessWithColor is a green string "Success: "
	SuccessWithColor = string(ColorGreen) + "Success: " + string(ColorWhite)
)

//InternalServerErrorMsg is an error that handles internal server error response
var InternalServerErrorMsg = echo.NewHTTPError(http.StatusInternalServerError, BaseError{Message: "internal_server_error"})
