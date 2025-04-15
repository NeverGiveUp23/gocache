# gocache

This is a project to better understamd L1-L3 cache lines while implementing padding for false sharing between cache lines, as well as simulating caches.

The first file 'cache.go' implements a fucntion to check spatial locality within my cache and how long thay process takes from a good spatial locality compared to bad spatial locality(cache thrashing).
-- Row-wise traversal accesses contiguous memory(exploiting cache lines), while column-wise jumps unpredictably (cache misses)
