package api

import (
	v1 "api-exam/api/handlers/v1"
	"api-exam/api/token"
	"api-exam/config"
	"api-exam/pkg/logger"
	"api-exam/services"
	"api-exam/storage/repo"

	"api-exam/api/middleware"

	"github.com/casbin/casbin/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
	Redis          repo.InMemorystorageI
	CasbinEnforcer *casbin.Enforcer
}

// @title  Exam API
// @version 1.0
// @description This is api for exam.

// @contact.url https://t.me/abdullohus

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option Option) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	jwtHandler := token.JWTHandler{
		SigninKey: option.Conf.SignKey,
		Log:       option.Logger,
	}

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
		Redis:          option.Redis,
		JWTHandler:     jwtHandler,
	})

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	router.Use(middleware.NewAuth(option.CasbinEnforcer, jwtHandler, config.Load()))

	api := router.Group("/v1")

	// customer

	api.POST("/register", handlerV1.Register)
	api.GET("/verify/:email/:code", handlerV1.Verification)
	api.GET("/login/:email/:password", handlerV1.Login)
	api.GET("/get-customer/:id", handlerV1.GetCustomerById)
	api.GET("/list-customer", handlerV1.GetCustomerList)
	api.DELETE("/delete-customer/:id", handlerV1.DeleteCustomerById)
	api.PUT("/update-customer", handlerV1.UpdateCustomer)

	// post
	api.POST("/create-post", handlerV1.CreatePost)
	api.PUT("/update-post", handlerV1.UpdatePost)
	api.GET("/get-post/:id", handlerV1.GetPost)
	api.GET("/list-post", handlerV1.GetPostList)
	api.DELETE("/delete-post/:id", handlerV1.DeletePostById)
	api.GET("/getpostlist/:id", handlerV1.GetPostByCustomerId)

	// reyting
	api.POST("/create-reyting", handlerV1.CreateReyting)

	// Admin
	api.GET("/admin-login/:email/:password", handlerV1.AdminLogin)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	return router
}
