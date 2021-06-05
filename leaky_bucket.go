package limiter

import "time"

type leaky struct {
	rate   int
	bucket chan struct{}
}

func NewLeakyBucket(rate int) Limiter {
	rl := leaky{
		rate:   rate,
		bucket: make(chan struct{}, rate),
	}
	for i := 0; i < rate; i++ {
		rl.bucket <- struct{}{}
	}
	go rl.start()
	return &rl
}

func (rl *leaky) Stop() {
	close(rl.bucket)
}

func (rl *leaky) Limit() {
	rl.bucket <- struct{}{}
}

func (rl *leaky) start() {
	ticker := time.NewTicker(time.Duration(1000/rl.rate) * time.Millisecond)
	defer ticker.Stop()
	for range ticker.C {
		_, keepLimiting := <-rl.bucket
		if !keepLimiting {
			return
		}
	}
}
