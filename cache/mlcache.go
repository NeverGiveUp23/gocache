package cache

type CacheSim struct {
	size       int // total size in bytes
	blockSize  int
	assoc      int
	accessTime int
	sets       []Set // Cache sets for N-Way associativity
}

type Set struct {
	blocks []Block
}

type Block struct {
	tag      int
	valid    bool
	lastUsed int // for LRU replacement
}

// Multi level cache hierarchy

type MultiLevelCache struct {
	l1         *CacheSim
	l2         *CacheSim
	l3         *CacheSim
	memLatency int // Ram access memLatency
	Stats      Stats
}

type Stats struct {
	L1Hits      int
	L2Hits      int
	L3Hits      int
	Misses      int
	TotalCycles int
}

func NewCacheSet(size, blockSize, assoc, accessTime int) *CacheSim {
	numSets := size / (blockSize * assoc)
	sets := make([]Set, numSets)

	for i := range sets {
		sets[i].blocks = make([]Block, assoc)
	}

	return &CacheSim{
		size:       size,
		blockSize:  blockSize,
		assoc:      assoc,
		accessTime: accessTime,
		sets:       sets,
	}
}

func NewMultilevelCache() *MultiLevelCache {
	return &MultiLevelCache{
		l1:         NewCacheSet(64, 16, 4, 1),
		l2:         NewCacheSet(256, 16, 4, 10),
		l3:         NewCacheSet(512, 16, 4, 30),
		memLatency: 100,
	}
}

// Simulate Memory Access

func (mlc *MultiLevelCache) Access(address int) {
	tagL1, indexL1 := getTagIndex(address, mlc.l1)
	tagL2, indexL2 := getTagIndex(address, mlc.l2)
	tagL3, indexL3 := getTagIndex(address, mlc.l3)

	// Check l1
	if mlc.checkCache(mlc.l1, tagL1, indexL1) {
		mlc.Stats.L1Hits++
		mlc.Stats.TotalCycles += mlc.l1.accessTime
		return
	}

	// check l2
	if mlc.checkCache(mlc.l2, tagL2, indexL2) {
		mlc.Stats.L2Hits++
		mlc.Stats.TotalCycles += mlc.l1.accessTime + mlc.l2.accessTime
		mlc.updateInclusiveCache(tagL1, indexL1, tagL2, indexL2) // bring into l1
		return
	}

	// check l3
	if mlc.checkCache(mlc.l3, tagL3, indexL3) {
		mlc.Stats.L3Hits++
		mlc.Stats.TotalCycles += mlc.l1.accessTime + mlc.l2.accessTime + mlc.l3.accessTime
		mlc.updateInclusiveCache(tagL1, indexL1, tagL2, indexL2) // bring into l1/l2
	}

	// Miss: Access RAM
	mlc.Stats.Misses++
	mlc.Stats.TotalCycles += mlc.l1.accessTime + mlc.l2.accessTime + mlc.l3.accessTime + mlc.memLatency
	mlc.updateInclusiveCache(tagL1, indexL1, tagL2, indexL2) // Brinv into l1,l2,l3
}

// Check if a block exists in a cache level
func (mlc *MultiLevelCache) checkCache(cache *CacheSim, tag, index int) bool {
	set := cache.sets[index]

	for _, block := range set.blocks {
		if block.valid && block.tag == tag {
			return true
		}
	}
	return false
}

// Add ablock to a cache
func (mlc *MultiLevelCache) addToCache(cache *CacheSim, tag, index int) {
	set := &cache.sets[index]
	for i := range set.blocks {
		if !set.blocks[i].valid {
			set.blocks[i] = Block{
				tag:      tag,
				valid:    true,
				lastUsed: 0,
			}
			return
		}
	}
	// Evict LRU Block
	lruIndex := 0
	for i := 1; i < cache.assoc; i++ {
		if set.blocks[i].lastUsed < set.blocks[lruIndex].lastUsed {
			lruIndex = 1
		}
	}

	set.blocks[lruIndex] = Block{
		tag:      tag,
		valid:    true,
		lastUsed: 0,
	}
}

func getTagIndex(address int, cache *CacheSim) (tag, index int) {
	blockAddr := address / cache.blockSize
	index = blockAddr % len(cache.sets)
	tag = blockAddr / len(cache.sets)
	return
}

func (mlc *MultiLevelCache) updateInclusiveCache(tagL1, indexL1, tagL2, indexL2 int) {
	mlc.addToCache(mlc.l1, tagL1, indexL1)
	mlc.addToCache(mlc.l2, tagL2, indexL2)
}
