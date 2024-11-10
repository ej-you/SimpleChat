package urls

import (
	echo "github.com/labstack/echo/v4"

	"SimpleChat/backend/app_chat/handlers"
	coreMiddlewares "SimpleChat/backend/core/middlewares"
)


func RouterGroup(group *echo.Group) {
	group.GET("/restricted", handlers.Restricted, coreMiddlewares.AuthMiddleware)
	group.GET("/free", handlers.Free)
}
