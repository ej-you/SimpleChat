package urls

import (
	echo "github.com/labstack/echo/v4"

	chatUrls "SimpleChat/backend/app_chat/urls"
	messangerUrls "SimpleChat/backend/app_messanger/urls"
	userUrls "SimpleChat/backend/app_user/urls"
)

// подгрузка urls каждого микроприложения и их общая настройка
func InitURLRouters(echoApp *echo.Echo) {
	apiUserGroup := echoApp.Group("/simple-chat/api/user")
	userUrls.RouterGroup(apiUserGroup)

	apiChatGroup := echoApp.Group("/simple-chat/api/chat")
	chatUrls.RouterGroup(apiChatGroup)

	apiMessangerGroup := echoApp.Group("/simple-chat/api/messanger")
	messangerUrls.RouterGroup(apiMessangerGroup)
}
