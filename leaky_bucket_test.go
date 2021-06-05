package limiter

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRateLimiter(t *testing.T) {
	rl := NewLeakyBucket(10)

	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 22; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			rl.Limit()
		}()
	}
	wg.Wait()
	rl.Stop()

	dur := time.Since(start)
	assert.True(t, dur > 2*time.Second, dur)
	assert.True(t, dur < 2*time.Second+210*time.Millisecond, dur)
}
