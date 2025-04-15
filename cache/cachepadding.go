package cache

import (
	"sync"
	"time"
)

// variable a and b in the same cache line 64 bytes
type NoPad struct {
	a int32
	b int32
}

type WithPad struct {
	a int32
	_ [60]byte // Pad to 64 bytes (common cache line size)
	b int32
}

func TestCounter(counter interface{}, times int) time.Duration {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 2; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for j := 0; j < times; j++ {
				// simulate concurrent writes
				switch c := counter.(type) {
				case *NoPad:
					c.a++
					c.b++
				case *WithPad:
					c.a++
					c.b++
				}
			}
		}()
	}
	wg.Wait()
	return time.Since(start)
}
