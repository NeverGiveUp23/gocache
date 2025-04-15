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

	cacheLine := cache.NewCache(64, 16)

	addresses := []int{0, 4, 16, 20, 32, 0, 16}

	for _, addr := range addresses {
		hit := cacheLine.SimulateAccess(addr)
		fmt.Printf("Address: %d: %v\n", addr, hit)
	}

	mlc := cache.NewMultilevelCache()

	addresses = []int{0x0, 0x4, 0x10, 0x14, 0x20, 0x0, 0x10}

	for _, addr := range addresses {
		mlc.Access(addr)
	}

	// Print statistics
	fmt.Printf("L1 Hits: %d\n", mlc.Stats.L1Hits)
	fmt.Printf("L2 Hits: %d\n", mlc.Stats.L2Hits)
	fmt.Printf("L3 Hits: %d\n", mlc.Stats.L3Hits)
	fmt.Printf("Misses: %d\n", mlc.Stats.Misses)
	fmt.Printf("Total Cycles: %d\n", mlc.Stats.TotalCycles)

}
