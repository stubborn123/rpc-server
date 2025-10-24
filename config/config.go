package config

import (
	"rpc-server/flags"
	"time"

	"github.com/urfave/cli/v2"
)

type Config struct {
	Migrations   string //迁移
	RpcServer    ServerConfig
	Metrics      ServerConfig
	RestServer   ServerConfig
	MasterDB     DBConfig
	SlaveDB      DBConfig
	BaseUrl      string
	LoopInternal time.Duration
}

// 服务器配置
type ServerConfig struct {
	Host string
	Port int
}

// 数据库配置
type DBConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
}

func NewConfig(ctx *cli.Context) Config {
	return Config{
		Migrations: ctx.String(flags.MigrationsFlag.Name),
	}
}
