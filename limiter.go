// Package limiter provides different types of rate limiters.
// Current types of rate limiters:
//  * Leaky Bucket (https://en.wikipedia.org/wiki/Leaky_bucket)
//
// Usage is as simple as:
//  rl := limiter.NewLeakyBucket(100)
//  defer rl.Stop()
//
//  for {
//    rl.Limit() // keep the process in check
//    ...
//  }
package limiter

type Limiter interface {
	// Stop will close the rate limiter
	Stop()

	// Limit should be called within the process or go-routines that should be rate limited.
	Limit()
}
