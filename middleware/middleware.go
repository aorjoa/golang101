package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// GateKeeper should logs all call to echo path
func GateKeeper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("something get call")
		return next(c)
	}
}