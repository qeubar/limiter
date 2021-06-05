# limiter

[![PkgGoDev](https://pkg.go.dev/badge/github.com/qeubar/limiter?tab=doc)](https://pkg.go.dev/github.com/qeubar/limiter?tab=doc)
[![Build Status](https://travis-ci.com/qeubar/limiter.svg?branch=master)](https://travis-ci.com/qeubar/limiter)

limiter is a simple and straight forward process rate limiter.

### Usage

```go

import "github.com/quebar/limiter"

func LimitedProcess() {
    rl := limiter.NewLeakyBucket(100)
    defer rl.Stop()

    var wg sync.WaitGroup
	for {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rl.Limit()
            // Do what you gotta do
		}()
	}
	wg.Wait()
}
