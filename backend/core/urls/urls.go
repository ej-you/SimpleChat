package urls

import (
	echo "github.com/labstack/echo/v4"

	chatUrls "SimpleChat/backend/app_chat/urls"
	messangerUrls "SimpleChat/backend/app_messanger/urls"
	userUrls "SimpleChat/backend/app_user/urls"
)

// подгрузка urls каждого микроприложения и их общая настройка
func InitURLRouters(echoApp *echo.Echo) {
	appGroup := echoApp.Group("/simple-chat/api")

	apiUserGroup := appGroup.Group("/user")
	userUrls.RouterGroup(apiUserGroup)

	apiChatGroup := appGroup.Group("/chat")
	chatUrls.RouterGroup(apiChatGroup)

	apiMessangerGroup := appGroup.Group("/messanger")
	messangerUrls.RouterGroup(apiMessangerGroup)
}
