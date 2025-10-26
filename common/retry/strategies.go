package retry

import (
	"math/rand"
	"time"
)

type Strategy interface {
	Duration(attempt int) time.Duration
}

type ExponentialStrategy struct {
	//time.Duration 时间间隔
	Min       time.Duration
	Max       time.Duration
	MaxJitter time.Duration
}

func (e *ExponentialStrategy) Duration(attempt int) time.Duration {
	var jitter time.Duration
	if e.MaxJitter > 0 {
		//Nanosecond 纳秒 10 -9秒
		jitter = time.Duration(rand.Int63n(e.MaxJitter.Nanoseconds()))
	}
	if attempt < 0 {
		return e.Min + jitter
	}

}
