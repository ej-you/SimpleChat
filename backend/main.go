package main

import (
	"fmt"
	"os"
	"time"

	echo "github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	// echoSwagger "github.com/swaggo/echo-swagger"
	// _ "SimpleChat/backend/docs"

	coreErrorHandler "SimpleChat/backend/core/error_handler"
	coreUrls "SimpleChat/backend/core/urls"

	"SimpleChat/backend/core/db"
	"SimpleChat/backend/settings"
)


// Настройка Swagger документации
//	@Title						Knofu Go API
//	@Version					1.0
//	@Description				This is a Knofu API written on Golang using Echo.
//	@Host						localhost:8000
//	@BasePath					/api
//	@Schemes					http
//	@Accept						json
//	@Produce					json
//	@SecurityDefinitions.apiKey	Access
//	@In							header
//	@Name						Authorization
//	@Description				JWT security accessToken. Please add it in the format "Bearer {AccessToken}" to authorize your requests.
//	@SecurityDefinitions.apiKey	Refresh
//	@In							header
//	@Name						Authorization
//	@Description				JWT security RefreshToken. Use it like "Bearer {RefreshToken}" to obtain new AccessToken.
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
		AllowCredentials: settings.CorsAllowCredentials,
	}))

	// настройка таймаута для всех запросов на 20 секунд
	echoApp.Use(echoMiddleware.TimeoutWithConfig(echoMiddleware.TimeoutConfig{
		ErrorMessage: "timeout error",
		Timeout: 20*time.Second,
	}))

	// настройка кастомного обработчика ошибок
	coreErrorHandler.CustomErrorHandler(echoApp)
	// настройка роутеров для эндпоинтов
	coreUrls.InitUrlRouters(echoApp)

	// натсройка Swagger документации
	// echoApp.GET("/api/swagger/*", echoSwagger.WrapHandler)

	// запуск приложения
	echoApp.Logger.Fatal(echoApp.Start(fmt.Sprintf(":%s", settings.Port)))
}
