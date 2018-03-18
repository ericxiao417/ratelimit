package main

import (
	"time"

	"github.com/ericxiao417/ratelimit"
)

func main() {
	r, err := ratelimit.NewThrottleRateLimiter(&ratelimit.Config{
		Throttle: 1 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	ratelimit.DoWork(r, 10)
}
