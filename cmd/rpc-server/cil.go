package main

import (
	"context"
	"rpc-server/common/cliapp"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/ethereum/go-ethereum/log"
	"github.com/urfave/cli/v2"
)

func runTask(ctx *cli.Context, shutdown context.CancelCauseFunc) (cliapp.Lifecycle, error) {
	log.Info("run market price task")
	cfg := config.NewConfig()
	db, err := database.NewDB()

}
