package main

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/the-web3/rpc-server/common/opio"
	"os"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	//这个log对应go常规的日志工具包，提供更多高级功能，如结构化日志，多级别日志控制，格式输出
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	app := NewCil()

	ctx := opio.WithInterruptBlocker(context.Background())

	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("Application failed")
		os.Exit(1)
	}

}
