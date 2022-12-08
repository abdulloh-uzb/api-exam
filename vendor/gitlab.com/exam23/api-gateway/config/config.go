package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment         string // develop, staging, production
	CustomerServiceHost string
	CustomerServicePort int
	RewiewServiceHost   string
	RewiewServicePort   int
	PostServiceHost     string
	PostServicePort     int
	CtxTimeout          int // context timeout in seconds
	LogLevel            string
	HTTPPort            string
	RedisHost           string
	RedisPort           string
	SignInKey           string
	PostgresHost        string
	PostgresPort        int
	PostgresUser        string
	PostgresDB          string
	PostgresPassword    string
	AuthConfigPath      string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(GetOrReturnDefault("ENVIRONMENT", "develop"))

	c.PostgresHost = cast.ToString(GetOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(GetOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDB = cast.ToString(GetOrReturnDefault("POSTGRES_DATABASE", "customer_service"))
	c.PostgresUser = cast.ToString(GetOrReturnDefault("POSTGRES_USER", "azizbek"))
	c.PostgresPassword = cast.ToString(GetOrReturnDefault("POSTGRES_PASSWORD", "Azizbek"))

	c.LogLevel = cast.ToString(GetOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(GetOrReturnDefault("HTTP_PORT", ":8080"))

	c.CustomerServiceHost = cast.ToString(GetOrReturnDefault("CUSTOMER_SERVICE_HOST", "127.0.0.1"))
	c.CustomerServicePort = cast.ToInt(GetOrReturnDefault("CUSTOMER_SERVICE_PORT", 7070))

	c.PostServiceHost = cast.ToString(GetOrReturnDefault("POST_SERVICE_HOST", "127.0.0.1"))
	c.PostServicePort = cast.ToInt(GetOrReturnDefault("POST_SERVICE_PORT", 9000))

	c.RewiewServiceHost = cast.ToString(GetOrReturnDefault("REWIEW_SERVICE_HOST", "127.0.0.1"))
	c.RewiewServicePort = cast.ToInt(GetOrReturnDefault("REWIEW_SERVICE_PORT", 8000))

	c.RedisHost = cast.ToString(GetOrReturnDefault("REDIS_HOST", "localhost"))
	c.RedisPort = cast.ToString(GetOrReturnDefault("REDIS_PORT", "6379"))

	c.SignInKey = cast.ToString(GetOrReturnDefault("SIGNINGKEY", "AzizbekSignIn"))
	c.CtxTimeout = cast.ToInt(GetOrReturnDefault("CTX_TIMEOUT", 7))
	c.AuthConfigPath = cast.ToString(GetOrReturnDefault("AUTH_PATH", "./config/auth.conf"))

	return c
}

func GetOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
