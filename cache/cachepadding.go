package cache

import (
	"sync"
	"time"
)

// Total 8 bytes (but padded to 64 bytes in cache lines)
type NoPad struct {
	a int32 // 4 bytes
	b int32 // 4 bytes
}

/*
Thread 1 updates a, Thread 2 updates b.
Even though their different variable, the entire cache line is invalidates on each write -> unnecessary synchronization
*/

// Solution is to add padding to ensure a and b reside in separate cache lines:
type WithPad struct {
	a int32
	_ [60]byte // Pad to 60 bytes
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
