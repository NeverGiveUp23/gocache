# gocache

This is a project to better understamd L1-L3 cache lines while implementing padding for false sharing between cache lines, as well as simulating caches.

The first file 'cache.go' implements a fucntion to check spatial locality within my cache and how long the process takes from a good spatial locality compared to bad spatial locality(cache thrashing).
-- Row-wise traversal accesses contiguous memory(exploiting cache lines), while column-wise jumps unpredictably (cache misses)

Test Case 2:

In cachepadding.go I wanted to test the time it takes from a false sharing data in the cache line vs padding implementation to avoid false sharing.

Thread 1 updates a, Thread 2 updates b.
Even though their different variable, the entire cache line is invalidates on each write -> unnecessary synchronization

Solution is to add padding to ensure a and b reside in separate cache lines.

    False sharing occurs when two or more threads (on different CPU cores) modify variables that reside on the same cache line, even though those variables are logically independent.

    Cache lines: Units of data transfer between CPU and memory (typically 64 bytes). When one thread writes to a cache line, other cores with the same line invalidate their copies, forcing a reload from memory.

    Performance impact: Frequent invalidations → cache misses → slower execution.

Test Case 3:

Involves getting cache hits/misses within our L1/L2/L3 cache and how many Cycles it took to complete

    L1 Hits: 2
    L2 Hits: 2
    L3 Hits: 0
    Misses: 3
    Total Cycles: 447
