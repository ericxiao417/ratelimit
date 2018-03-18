package main

import (
	"time"

	"github.com/ericxiao417/ratelimit"
)

func main() {
	r, err := ratelimit.NewFixedWindowRateLimiter(&ratelimit.Config{
		Limit:         5,
		FixedInterval: 15 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	ratelimit.DoWork(r, 10)
}
