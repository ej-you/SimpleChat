package urls

import (
	echo "github.com/labstack/echo/v4"

	"SimpleChat/backend/app_messanger/handlers"
	coreMiddlewares "SimpleChat/backend/core/middlewares"
)


func RouterGroup(group *echo.Group) {
	group.GET("", handlers.UpgradeWebSocket, coreMiddlewares.AuthMiddleware)
}
