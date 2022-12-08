package config

import (
	"os"

	"github.com/spf13/cast"
)

// Config ...
type Config struct {
	Environment string // develop, staging, production

	CustomerServiceHost string
	CustomerServicePort int
	ReytingServiceHost  string
	ReytingServicePort  int
	PostServiceHost     string
	PostServicePort     int

	// context timeout in seconds
	CtxTimeout       int
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	AuthConfigPath   string
	SignKey          string
	RedisHost        string
	RedisPort        string
	LogLevel         string
	HTTPPort         string
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))

	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))

	c.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	c.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	c.PostgresDB = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "user_1"))
	c.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	c.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "1234"))
	c.RedisHost = cast.ToString(getOrReturnDefault("REDIS_HOST", "localhost"))
	c.RedisPort = cast.ToString(getOrReturnDefault("REDIS_PORT", "6379"))

	c.AuthConfigPath = cast.ToString(getOrReturnDefault("AUTH_CONFIG_PATH", "./config/auth.conf"))
	c.SignKey = cast.ToString(getOrReturnDefault("SIGNIN_KEY", "usmonov"))

	c.CustomerServiceHost = cast.ToString(getOrReturnDefault("CUSTOMER_SERVICE_HOST", "127.0.0.1"))
	c.CustomerServicePort = cast.ToInt(getOrReturnDefault("CUSTOMER_SERVICE_PORT", 8002))
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "127.0.0.1"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 8000))
	c.ReytingServiceHost = cast.ToString(getOrReturnDefault("REYTING_SERVICE_HOST", "127.0.0.1"))
	c.ReytingServicePort = cast.ToInt(getOrReturnDefault("REYTING_SERVICE_PORT", 1111))

	c.CtxTimeout = cast.ToInt(getOrReturnDefault("CTX_TIMEOUT", 7))

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
