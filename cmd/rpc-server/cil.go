package main

import (
	"context"
	"rpc-server/common/cliapp"
	"rpc-server/config"

	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

func runTask(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("run market price task")

	//这里的NewConfig 和 NewDB可以看出和Java的特性对比，Java是面向对象，创建是直接new一个class定义的指定类型的实例
	//Go是面向组合的，COP---用结构体和接口，通过组合实现复用（这里和Java不同，而是直接去对应的包里面获取对应公共函数创建实例）
	cfg := config.NewConfig(ctx)
	db, err := database.NewDB()

}
