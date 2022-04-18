package framework

import (
	"fmt"
	"go.uber.org/fx"
	"kloudlite.io/apps/finance/internal/app"
	"kloudlite.io/pkg/cache"
	"kloudlite.io/pkg/config"
	rpc "kloudlite.io/pkg/grpc"
	httpServer "kloudlite.io/pkg/http-server"
	"kloudlite.io/pkg/logger"
	"kloudlite.io/pkg/repos"
)

type Env struct {
	DBName        string `env:"MONGO_DB_NAME"`
	DBUrl         string `env:"MONGO_URI"`
	RedisHosts    string `env:"REDIS_HOSTS"`
	RedisUsername string `env:"REDIS_USERNAME"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	HttpPort      uint16 `env:"PORT"`
	HttpCors      string `env:"ORIGINS"`
	IAMServerHost string `env:"IAM_SERVER_HOST"`
	IAMServerPort uint16 `env:"IAM_SERVER_PORT"`
}

func (e *Env) GetMongoConfig() (url string, dbName string) {
	return e.DBUrl, e.DBName
}

func (e *Env) RedisOptions() (hosts, username, password string) {
	return e.RedisHosts, e.RedisUsername, e.RedisPassword
}

func (e *Env) GetHttpPort() uint16 {
	return e.HttpPort
}

func (e *Env) GetHttpCors() string {
	return e.HttpCors
}

func (e *Env) GetGCPServerURL() string {
	return fmt.Sprintf("%v:%v", e.IAMServerHost, e.IAMServerPort)
}

var Module = fx.Module("framework",
	fx.Provide(logger.NewLogger),
	config.EnvFx[Env](),
	rpc.NewGrpcClientFx[*Env](),
	repos.NewMongoClientFx[*Env](),
	cache.NewRedisFx[*Env](),
	httpServer.NewHttpServerFx[*Env](),
	app.Module,
)
