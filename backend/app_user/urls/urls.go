package urls

import (
	echo "github.com/labstack/echo/v4"

	"SimpleChat/backend/app_user/handlers"
	coreMiddlewares "SimpleChat/backend/core/middlewares"
)


func RouterGroup(group *echo.Group) {
	group.POST("/register", handlers.Register)
	group.POST("/login", handlers.Login)

	group.GET("/check/:username", handlers.Check, coreMiddlewares.AuthMiddleware)
}
