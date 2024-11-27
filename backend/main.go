package main

import (
	"fmt"
	"os"
	"time"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
	_ "SimpleChat/backend/docs"

	coreErrorHandler "SimpleChat/backend/core/error_handler"
	coreUrls "SimpleChat/backend/core/urls"

	"SimpleChat/backend/core/db"
	"SimpleChat/backend/settings"
)


// Настройка Swagger документации
//	@Title						SimpleChat Go API
//	@Version					1.0
//	@Description				This is a SimpleChat API written on Golang using Echo and Gorilla WebSocket.
//	@Host						127.0.0.1:8000
//	@Host						150.241.82.68
//	@BasePath					/api
//	@Schemes					http
//	@Accept						json
//	@Produce					json
//	@SecurityDefinitions.apiKey	CookieAuth
//	@In							cookie
//	@Name						auth
//	@Description				JWT security token. Cookie is automatic added after auth is done (login/register).
func main() {
	echoApp := echo.New()
	echoApp.HideBanner = true

	// если при запуске указан аргумент "dev" или "migrate"
	args := os.Args
	if len(args) > 1 {
		// запуск в dev режиме
		if args[1] == "dev" {
			echoApp.Debug = true
		// проведение миграций БД без запуска самого приложения
		} else if args[1] == "migrate" {
			db.Migrate()
			return
		}
	}

	// удаление последнего слеша
	echoApp.Pre(echoMiddleware.RemoveTrailingSlash())
	// кастомизация логирования
	echoApp.Use(echoMiddleware.LoggerWithConfig(echoMiddleware.LoggerConfig{
		Format: settings.LogFmt,
	}))
	// отлавливание паник для беспрерывной работы сервиса
	echoApp.Use(echoMiddleware.Recover())

	// настройка CORS
	echoApp.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: settings.CorsAllowedOrigins,
		AllowMethods: settings.CorsAllowedMethods,
		AllowHeaders: []string{"Content-Type", "Authorization", "Upgrade", "Sec-WebSocket-Protocol", "Sec-WebSocket-Key", "Sec-WebSocket-Version"},
		AllowCredentials: true,
	}))

	// настройка таймаута для всех HTTP запросов на 20 секунд
	echoApp.Use(echoMiddleware.TimeoutWithConfig(echoMiddleware.TimeoutConfig{
		// пропускаем использование этого middleware для WebSocket соединения 
		Skipper: func(context echo.Context) bool {
			if context.Request().URL.Path == settings.WebsocketURLPath {
				return true
			}
			return false
		},
		ErrorMessage: "timeout error",
		Timeout: 20*time.Second,
	}))

	// настройка кастомного обработчика ошибок
	coreErrorHandler.CustomErrorHandler(echoApp)
	// настройка роутеров для эндпоинтов
	coreUrls.InitUrlRouters(echoApp)

	// настройка Swagger документации
	echoApp.GET("/api/swagger/*", echoSwagger.WrapHandler)

	// запуск приложения
	echoApp.Logger.Fatal(echoApp.Start(fmt.Sprintf(":%s", settings.Port)))
}
