# ratelimit

[![GoDoc](https://godoc.org/github.com/ericxiao417/ratelimit?status.svg)](https://godoc.org/github.com/ericxiao417/ratelimit)

`ratelimit` is a Go package that provides common rate limiting algorithms. It supports various rate limiting strategies including fixed window, max concurrency, and throttling.

## Features

- **Fixed Window Rate Limiter**: Limits the number of requests in a fixed time window.
- **Max Concurrency Rate Limiter**: Limits the number of concurrent requests.
- **Throttle Rate Limiter**: Ensures a minimum time interval between requests.

## Installation

To install the package, run:

```sh
go get github.com/ericxiao417/ratelimit
```

## Usage

### Fixed Window Rate Limiter

The Fixed Window Rate Limiter limits the number of requests in a fixed time window.

```go
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
```

### Max Concurrency Rate Limiter

The Max Concurrency Rate Limiter limits the number of concurrent requests.

```go
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
```

### Throttle Rate Limiter

The Throttle Rate Limiter ensures a minimum time interval between requests.

```go
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
```

## Configuration

The `Config` struct is used to configure the rate limiters:

```go
type Config struct {
	// Limit determines how many rate limit tokens can be active at a time
	Limit int

	// FixedInterval sets the fixed time window for a Fixed Window Rate Limiter
	FixedInterval time.Duration

	// Throttle is the min time between requests for a Throttle Rate Limiter
	Throttle time.Duration

	// TokenResetsAfter is the maximum amount of time a token can live before being
	// forcefully released - if set to zero time then the token may live forever
	TokenResetsAfter time.Duration
}
```

## Error Handling

The package defines several errors that can be used for error handling:

```go
var (
	ErrInvalidLimit           = errors.New("Limit must be greater than zero")
	ErrInvalidInterval        = errors.New("Interval must be greater than zero")
	ErrTokenFactoryNotDefined = errors.New("Token factory must be defined")
)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Acknowledgements

This package was created by [ericxiao417](https://github.com/ericxiao417).

## References

- [GoDoc](https://godoc.org/github.com/ericxiao417/ratelimit)
```
