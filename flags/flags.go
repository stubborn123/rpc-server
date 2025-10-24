package flags

import "github.com/urfave/cli/v2"

//定义Flag标识

// 定义一个不可改变的常量envVarPrefix
const envVarPrefix = "MARKET"

// 作用就是返回一个字符串数组，一个标识带“MARKET”的环境
func prefixEnvVars(name string) []string {
	return []string{envVarPrefix + "_" + name}
}

var (
	//StringFlag是urfave/cli/v2库中用于定义命令行字符串类型参数的结构体
	//采用cli/v2的结构体，而不是自己单独自己定义，因为项目使用了urfave/cli/v2作为命令行应用框架
	//使用这个StrinfFlag，库中预定义是一个标准的结构体，另一方面对于库的集成（如参数解析，环境变量绑定等，自定义的结构体无法直接利用这些功能）

	MigrationsFlag = &cli.StringFlag{
		Name:    "migrations-dir",
		Value:   "./migrations",
		Usage:   "path for database migrations",
		EnvVars: prefixEnvVars("MIGRATIONS_DIR"),
	}

	//RPC service (host , port)
	RpcHostFlag = &cli.StringSliceFlag{
		Name:     "rpc-host",
		Usage:    "The host of the rpc",
		EnvVars:  prefixEnvVars("RPC_HOST"),
		Required: true,
	}

	RpcPortFlag = &cli.StringSliceFlag{
		Name:     "rpc-port",
		Usage:    "The port of the rpc",
		EnvVars:  prefixEnvVars("RPC_PORT"),
		Required: true,
	}

	HttpHostFlag = &cli.StringSliceFlag{
		Name:     "http-host",
		Usage:    "The host of the http",
		EnvVars:  prefixEnvVars("HTTP_HOST"),
		Required: true,
	}

	HttpPortFlag = &cli.StringSliceFlag{
		Name:     "http-port",
		Usage:    "THe port of the http",
		EnvVars:  prefixEnvVars("HTTP_PORT"),
		Required: true,
	}
	// MasterDbHostFlag Flags
	MasterDbHostFlag = &cli.StringFlag{
		Name:     "master-db-host",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_HOST"),
		Required: true,
	}
	MasterDbPortFlag = &cli.IntFlag{
		Name:     "master-db-port",
		Usage:    "The port of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_PORT"),
		Required: true,
	}
	MasterDbUserFlag = &cli.StringFlag{
		Name:     "master-db-user",
		Usage:    "The user of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_USER"),
		Required: true,
	}
	MasterDbPasswordFlag = &cli.StringFlag{
		Name:     "master-db-password",
		Usage:    "The host of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_PASSWORD"),
		Required: true,
	}
	MasterDbNameFlag = &cli.StringFlag{
		Name:     "master-db-name",
		Usage:    "The db name of the master database",
		EnvVars:  prefixEnvVars("MASTER_DB_NAME"),
		Required: true,
	}
	LoopInternalFlag = &cli.DurationFlag{
		Name:     "loop-internal",
		Usage:    "task exec cycle",
		EnvVars:  prefixEnvVars("LOOP_INTERNAL"),
		Required: true,
	}
	BaseUrlFlag = &cli.StringFlag{
		Name:     "base-url",
		Usage:    "Base url of exchange",
		EnvVars:  prefixEnvVars("BASE_URL"),
		Required: true,
	}

	// Slave DB  flags
	SlaveDbHostFlag = &cli.StringFlag{
		Name:    "slave-db-host",
		Usage:   "The host of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_HOST"),
	}
	SlaveDbPortFlag = &cli.IntFlag{
		Name:    "slave-db-port",
		Usage:   "The port of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_PORT"),
	}
	SlaveDbUserFlag = &cli.StringFlag{
		Name:    "slave-db-user",
		Usage:   "The user of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_USER"),
	}
	SlaveDbPasswordFlag = &cli.StringFlag{
		Name:    "slave-db-password",
		Usage:   "The host of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_PASSWORD"),
	}
	SlaveDbNameFlag = &cli.StringFlag{
		Name:    "slave-db-name",
		Usage:   "The db name of the slave database",
		EnvVars: prefixEnvVars("SLAVE_DB_NAME"),
	}

	MetricsHostFlag = &cli.StringFlag{
		Name:     "metric-host",
		Usage:    "The host of the metric",
		EnvVars:  prefixEnvVars("METRIC_HOST"),
		Required: true,
	}
	MetricsPortFlag = &cli.IntFlag{
		Name:     "metric-port",
		Usage:    "The port of the metric",
		EnvVars:  prefixEnvVars("METRIC_PORT"),
		Required: true,
	}
)

var requireFlags = []cli.Flag{
	MigrationsFlag,
	RpcHostFlag,
	RpcPortFlag,
	HttpHostFlag,
	HttpPortFlag,
	MasterDbHostFlag,
	MasterDbPortFlag,
	MasterDbUserFlag,
	MasterDbPasswordFlag,
	MasterDbNameFlag,
	LoopInternalFlag,
	BaseUrlFlag,
}

var optionalFlags = []cli.Flag{
	SlaveDbHostFlag,
	SlaveDbPortFlag,
	SlaveDbUserFlag,
	SlaveDbPasswordFlag,
	SlaveDbNameFlag,
	MetricsHostFlag,
	MetricsPortFlag,
}

func init() {
	Flags = append(requireFlags, optionalFlags...)
}

var Flags []cli.Flag
