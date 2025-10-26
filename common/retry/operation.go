package retry

import (
	"context"
	"fmt"
)

// 错误重试设置
type ErrFailedPermanently struct {
	attempts int
	LastErr  error
}

func (e *ErrFailedPermanently) Error() string {
	//Sprintf对比Printf，有返回结果更适配。Printf就是直接在控制台打印
	return fmt.Sprintf("operation failed permanently after:%v", e.attempts, e.LastErr)
}

func (e *ErrFailedPermanently) Unwrap() error {
	return e.LastErr
}

type pair[T, U any] struct {
	a T
	b U
}

func Do2[T, U any](ctx context.Context, maxAttempts int, strategy Stragetgy, op func() (T, U, error)) (T, U, error) {
	f := func() (pair[T, U], error) {
		a, b, err := op()
		return pair[T, U]{a, b}, err
	}
	res.err := Do(ctx, maxAttempts, strategy, f)
	return res.a, res.b.err
}

func Do[T, U any](ctx context.Context, maxAttempts int, strategy Strategy, op func() (T, U, error)) {

}
