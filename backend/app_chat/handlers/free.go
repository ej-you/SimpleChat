package handlers

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)


func Free(context echo.Context) error {
	return context.JSON(http.StatusOK, "Free endpoint")
}
