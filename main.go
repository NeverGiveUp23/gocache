package main

import (
	"fmt"
	"github.com/nevergiveup23/gocacheproject/cache"
	"time"
)

const (
	Rows = 1024
	Cols = 1024
)

func main() {
	// initialize a large matrix
	matrix := make([][]int, Rows)
	for i := range matrix {
		matrix[i] = make([]int, Cols)
	}

	// Time row-wise traversal
	start := time.Now()
	cache.RowWiseTraversal(matrix)
	rowTime := time.Since(start)
	// Time column-wise traversal
	start = time.Now()
	cache.ColumTraversal(matrix)
	colsTime := time.Since(start)

	fmt.Printf("Row-wise: %v\nColumn-wise: %v\n", rowTime, colsTime)

	// output:
	// Row-wise: 618.959µs
	// Column-wise: 769.041µs
	/*
	   Row-wise traversal accesses contiguous memory(exploiting cache lines), while column-wise jumps unpredictably (cache misses)
	*/

	noPad := &cache.NoPad{}
	withPad := &cache.WithPad{}

	iterations := 1_000_000

	fmt.Println("No Padding:", cache.TestCounter(noPad, iterations))
	fmt.Println("With Padding:", cache.TestCounter(withPad, iterations))

	/*
	   No Padding: 9.4ms
	   With Padding: 5.6ms
	*/

	cache := cache.NewCache(64, 16)

	addresses := []int{0, 4, 16, 20, 32, 0, 16}

	for _, addr := range addresses {
		hit := cache.SimulateAccess(addr)
		fmt.Printf("Address: %d: %v\n", addr, hit)
	}

}
