package opio

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// 定义一个操作系统信号的切片，包含了四种常见的终止/中断信号（os和请求进程）
var DefaultInterruptSignals = []os.Signal{
	os.Interrupt,
	os.Kill,
	syscall.SIGTERM,
	syscall.SIGQUIT,
}

// 入参：接收可变数量的操作系统信号参数
func BlockOnInterrupts(signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultInterruptSignals
	}
	//创建一个容量为1的channel，接收信号变量
	interruptChannel := make(chan os.Signal, 1)
	//Notify方法，入参channel ，可变参数signals
	//这里的channel是引用类型，不必像Go普通变量想要被操作修改需要传递指针，channel可以被Notify方法修改
	//这个channel使用和Java的传参很像，不像go其他变量如果不用指针只能传递一个副本，导致这个原变量无法被修改
	signal.Notify(interruptChannel, signals...)

	//这个代码你可以理解，channel是一个mq，signal的通知方法是一个生产者向管道添加数据
	//这个管道会消费也可以叫接收数据
	<-interruptChannel
}

func BlockOnInterruptsXContext(ctx context.Context, signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultInterruptSignals
	}
	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, signals...)
	select {
	case <-interruptChannel:
	case <-ctx.Done():
		signal.Stop(interruptChannel)
	}
}

type interruptContextKeyType struct {
}

var blockerContextKey = interruptContextKeyType{}

type interruptCatcher struct {
	incoming chan os.Signal
}

func (c *interruptCatcher) Block(ctx context.Context) {
	select {
	case <-c.incoming:
	case <-ctx.Done():
	}
}

func WithInterruptBlocker(ctx context.Context) context.Context {
	if ctx.Value(blockerContextKey) != nil {
		return ctx
	}
	catcher := &interruptCatcher{
		incoming: make(chan os.Signal, 10),
	}
	signal.Notify(catcher.incoming, DefaultInterruptSignals...)
	return context.WithValue(ctx, blockerContextKey, BlockFn(catcher.Block))
}

func WithBlock(ctx context.Context, fn BlockFn) context.Context {
	return context.WithValue(ctx, blockerContextKey, fn)
}

type BlockFn func(ctx context.Context)

func BlockFromContext(ctx context.Context) BlockFn {
	v := ctx.Value(blockerContextKey)
	if v == nil {
		return nil
	}
	return v.(BlockFn)
}

func CancelOnInterrupt(ctx context.Context) context.Context {
	inner, cancel := context.WithCancel(ctx)

	blockOnInterrupt := BlockFromContext(ctx)

	if blockOnInterrupt == nil {
		blockOnInterrupt = func(ctx context.Context) {
			BlockOnInterruptsXContext(ctx)
		}
	}

	go func() {
		blockOnInterrupt(ctx)
		cancel()
	}()

	return inner
}
