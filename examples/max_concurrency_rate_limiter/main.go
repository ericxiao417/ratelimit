package main

import (
	"time"

	"github.com/ericxiao417/ratelimit"
)

func main() {
	r, err := ratelimit.NewMaxConcurrencyRateLimiter(&ratelimit.Config{
		Limit:            4,
		TokenResetsAfter: 10 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	ratelimit.DoWork(r, 10)
}
