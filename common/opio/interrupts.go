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
	//interruptChannel没有数据就会阻塞
	//interruptChannel没有数据阻塞状态会接触，程序执行后续代码
	<-interruptChannel
}

// 这是一个增强的信号监听函数（对比前面那个不带Context的函数）
func BlockOnInterruptsContext(ctx context.Context, signals ...os.Signal) {
	if len(signals) == 0 {
		signals = DefaultInterruptSignals
	}
	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, signals...)

	//和上面的监听函数不同这里添加了select
	//通过select管道操作的关键字，类似与Unix的select系统调用，可以同时监控多个通道的读写状态
	select {
	case <-interruptChannel:
		//当接收到信号时执行
	case <-ctx.Done():
		//当上下文取消时执行
		signal.Stop(interruptChannel)
	}
}

// type关键字对应的是声明结构体，var是对应的是变量
type interruptContextKeyType struct {
}

// 实际上结合上面的代码，两句代码结合成一行代码 var blockerContextKey = struct{}{},但是不推荐，有类型冲突风险，其他代码也可以定义一个完全相同的空结构体，导致上下文键冲突
// 同样结合上面代码，上面的type是一个类型，这个interruptContextKeyType{}是一个实例，类似与Java定义一个class，在创建一个实例
var blockerContextKey = interruptContextKeyType{}

type interruptCatcher struct {
	incoming chan os.Signal
}

// 这个 c *interruptCather 指针，是一个接收者receiver，因为go的函数参数是副本这种，为例面向对象就创造一个接收者
// 这个receiver是go面向对象的一种方式，他表示的是定义这个函数属于某个类型。
// 这个接收者的调用方法使用和普通的函数的调用方式不一样，要同住指定类型的实例来调用（和Java的对象实例调用方法很像）
// 具体调用使用可以看一下WithInterruptBlocker方法
func (c *interruptCatcher) Block(ctx context.Context) {

	//go的select很有意思就是，当多个管道都有数据，也就是满足多个case，go会随机选择一个case执行
	//好处是避免饥饿，防止永远是一个case被执行别的case不执行，坏处是随机的增加了不确定性
	select {
	case <-c.incoming:
		//incoming有数据执行
	case <-ctx.Done():
		//ctx上下文取消执行
	}
}

func WithInterruptBlocker(ctx context.Context) context.Context {
	//首先看context的源码，了解go的接口定义，和Java不同，虽然也是单独的一个关键字interface，但是不在是一个单独的文件
	//一个interface关键字，在它的结构体里的方法就是定义的接口方法，比如这个Value方法在context里面是Value(key any) any

	//这里表示ctx存储了这个键，有就返回，没有就返回nil
	if ctx.Value(blockerContextKey) != nil {
		return ctx
	}
	//创建一个实例的指针，方便共享给不同的结构体
	catcher := &interruptCatcher{
		incoming: make(chan os.Signal, 10),
	}
	signal.Notify(catcher.incoming, DefaultInterruptSignals...)
	//withValue是context的一个函数，主要作用是用于为上下文添加键值对
	// catcher.Block类似实例调用方法
	// BlockFn(cather.Block)是一个函数类型实例（应为是函数，不需要像struct结构体的实例还要添加一个{}）
	return context.WithValue(ctx, blockerContextKey, BlockFn(catcher.Block))
}

func WithBlock(ctx context.Context, fn BlockFn) context.Context {
	return context.WithValue(ctx, blockerContextKey, fn)
}

type BlockFn func(ctx context.Context)

func BlockFromContext(ctx context.Context) BlockFn {

	//因为Value返回的是any，any等价于interface{}空接口，可以表示是任意类型的值
	v := ctx.Value(blockerContextKey)
	if v == nil {
		return nil
	}
	// v.(BlockFn)是一个类型断言，检查一个接口类型的值是否为特定的具体类型
	// 如果转换成功，则返回一个BlockFn
	// 转换不成功，第二个参数ok变为false，可以根据这个来判断
	return v.(BlockFn)
}

func CancelOnInterrupt(ctx context.Context) context.Context {
	//看源码：inner是一个上下文，cancel是一个函数实例
	inner, cancel := context.WithCancel(ctx)

	blockOnInterrupt := BlockFromContext(ctx)

	if blockOnInterrupt == nil {
		blockOnInterrupt = func(ctx context.Context) {
			//监听是否中断
			BlockOnInterruptsContext(ctx)
		}
	}

	//在后台启动一个协程，监听中断信号并在收到信号时取消上下文。
	go func() {
		//在这个 goroutine 中，首先调用 blockOnInterrupt(ctx)，这个函数会阻塞直到收到中断信号或上下文被取消
		blockOnInterrupt(ctx)
		//当 blockOnInterrupt 函数返回时（通常是收到了中断信号），立即调用 cancel() 函数取消 inner 上下文
		cancel()
	}()

	return inner
}
