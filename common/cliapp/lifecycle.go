package cliapp

import (
	"context"
	"errors"
	"fmt"
	"os"
	"rpc-server/common/opio"

	"github.com/urfave/cli/v2"
)

// 首字母大写，提供给外面（go里面的接口也可以作为参数）
type Lifecycle interface {

	//注解方法首字母大写，相当于public修饰符
	Start(ctx context.Context) error

	Stop(ctx context.Context) error

	Stopped() bool
}

type LifecycleAction func(ctx *cli.Context, close context.CancelCauseFunc) (Lifecycle, error)

// cli.ActionFunc 用来定义命令行命令的执行逻辑
func LifecycleCmd(fn LifecycleAction) cli.ActionFunc {
	return lifecycleCmd(fn, opio.BlockOnInterruptsContext)
}

type waitSignalFn func(ctx context.Context, signals ...os.Signal)

var interruptErr = errors.New("interrupt signal")

func lifecycleCmd(fn LifecycleAction, blockOnInterrupt waitSignalFn) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		hostCtx := ctx.Context
		appCtx, appCancel := context.WithCancelCause(hostCtx)

		ctx.Context = appCtx

		go func() {
			blockOnInterrupt(appCtx)
			appCancel(interruptErr)
		}()

		appLifecycle, err := fn(ctx, appCancel)

		if err != nil {
			return errors.Join(
				//%w 是一个包装错误，保留原始引用
				fmt.Errorf("failed to setup: %w", err),
				context.Cause(appCtx),
			)
		}

		<-appCtx.Done()

		stopCtx, stopCancel := context.WithCancelCause(hostCtx)
		go func() {
			blockOnInterrupt(stopCtx)
			stopCancel(interruptErr)
		}()

		stopErr := appLifecycle.Stop(stopCtx)
		stopCancel(nil)

		if stopErr != nil {
			return errors.Join(
				fmt.Errorf("failed to stop: %w", err),
				context.Cause(stopCtx),
			)
		}

		return nil
	}

}
