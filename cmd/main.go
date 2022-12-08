package main

import (
	"api-exam/api"
	_ "api-exam/api/docs"
	"api-exam/config"
	"api-exam/pkg/logger"
	"api-exam/services"

	// gormadapter "github.com/casbin/gorm-adapter/v2"
	r "api-exam/storage/redis"

	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	defaultrolemanager "github.com/casbin/casbin/v2/rbac/default-role-manager"
	"github.com/casbin/casbin/v2/util"

	"github.com/gomodule/redigo/redis"
)

func main() {
	var (
		casbinEnforcer *casbin.Enforcer
	)
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "api_gateway")

	rules := fileadapter.NewAdapter("./config/policy.csv")

	casbinEnforcer, err := casbin.NewEnforcer(cfg.AuthConfigPath, rules)

	if err != nil {

		log.Error("casbin enforcer error", logger.Error(err))

		return

	}

	err = casbinEnforcer.LoadPolicy()

	if err != nil {
		log.Error("error while load policy")
		return
	}

	serviceManager, err := services.NewServiceManager(&cfg)

	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
		return
	}

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch", util.KeyMatch)

	casbinEnforcer.GetRoleManager().(*defaultrolemanager.RoleManager).AddMatchingFunc("keyMatch3", util.KeyMatch3)

	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
		return
	}

	rdb := &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", cfg.RedisHost+":"+cfg.RedisPort)
		},
	}

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
		Redis:          r.NewRedisRepo(rdb),
		CasbinEnforcer: casbinEnforcer,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server",
			logger.Error(err))
		panic(err)
	}
}
