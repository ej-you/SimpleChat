package handlers

import (
	// "fmt"
	"net/http"

	echo "github.com/labstack/echo/v4"
)


func Restricted(context echo.Context) error {
	return context.JSON(http.StatusOK, "Restricted endpoint")
}
