package urls

import (
	echo "github.com/labstack/echo/v4"

	"SimpleChat/backend/app_user/handlers"
)


func RouterGroup(group *echo.Group) {
	group.POST("/register", handlers.Register)
	group.POST("/login", handlers.Login)
}
