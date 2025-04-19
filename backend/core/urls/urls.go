package urls

import (
	echo "github.com/labstack/echo/v4"

	chatUrls "SimpleChat/backend/app_chat/urls"
	messangerUrls "SimpleChat/backend/app_messanger/urls"
	userUrls "SimpleChat/backend/app_user/urls"
)

// подгрузка urls каждого микроприложения и их общая настройка
func InitUrlRouters(echoApp *echo.Echo) {
	apiUserGroup := echoApp.Group("/api/user")
	userUrls.RouterGroup(apiUserGroup)

	apiChatGroup := echoApp.Group("/api/chat")
	chatUrls.RouterGroup(apiChatGroup)

	apiMessangerGroup := echoApp.Group("/api/messanger")
	messangerUrls.RouterGroup(apiMessangerGroup)
}
