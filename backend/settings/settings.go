package settings

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

// загрузка переменных окружения
var _ error = godotenv.Load("./.env")

// распаковка переменных окружения
var Port string = os.Getenv("GO_PORT")
var SecretForJWT string = os.Getenv("SECRET")

const WebsocketURLPath = "/simple-chat/api/messanger"
const WebsocketPongWait = time.Second * 60
const WebsocketPingPeriod = WebsocketPongWait * 9 / 10

// настройки CORS
var CorsAllowedOrigins []string = strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
var CorsAllowedMethods []string = strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")

// параметры для настройки куки авторизации
var CookieSecure = func() bool {
	cookieSecureValue := os.Getenv("COOKIES_SECURE")
	if cookieSecureValue == "true" {
		return true
	}
	return false
}()
var CookieSameSite http.SameSite = func() http.SameSite {
	sameSiteValue := os.Getenv("COOKIES_SAME_SITE")

	switch sameSiteValue {
	case "LaxMode":
		return http.SameSiteLaxMode
	case "StrictMode":
		return http.SameSiteStrictMode
	case "NoneMode":
		return http.SameSiteNoneMode
	default:
		return http.SameSiteDefaultMode
	}
}()

// время истечения действия токена
var TokenExpiredTime time.Duration = time.Hour * 720

// путь до SQLite3 БД - os.Getenv("PATH_DB") || "./db.sqlite3"
var PathDB = func() string {
	dbPathEnv, isExists := os.LookupEnv("PATH_DB")
	// если перменная окружения не указана
	if !isExists {
		return "./db.sqlite3"
	}
	return dbPathEnv
}()

// формат логов
var LogFmt = "[${time_rfc3339}] -- ${status} -- from ${remote_ip} to ${host} (${method} ${uri}) [time: ${latency_human}] | ${bytes_in} ${bytes_out} | error: ${error} | -> User-Agent: ${user_agent}\n"

// формат времени
var TimeFmt = "06-01-02 15:04:05 -07"

// логеры
var InfoLog *log.Logger = log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
var WarnLog *log.Logger = log.New(os.Stdout, "[WARN]\t", log.Ldate|log.Ltime|log.Lshortfile)
var ErrorLog *log.Logger = log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

// функция для обработки критических ошибок
func DieIf(err error) {
	if err != nil {
		panic(err)
	}
}
