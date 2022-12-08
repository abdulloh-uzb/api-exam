package v1

import (
	t "api-exam/api/token"
	"api-exam/config"
	"api-exam/pkg/logger"
	"api-exam/services"
	"api-exam/storage/repo"
)

type handlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
	redis          repo.InMemorystorageI
	jwthandler     t.JWTHandler
}

// handlerV2Config ...

type HandlerV1Config struct {
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Cfg            config.Config
	Redis          repo.InMemorystorageI
	JWTHandler     t.JWTHandler
}

// New ...

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:            c.Logger,
		serviceManager: c.ServiceManager,
		cfg:            c.Cfg,
		redis:          c.Redis,
		jwthandler:     c.JWTHandler,
	}

}
