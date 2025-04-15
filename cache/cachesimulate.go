package cache

type Cache struct {
	numLines  int
	blockSize int
	tags      []int
	valid     []bool
}

func NewCache(totalSize, blockSize int) *Cache {
	numLines := totalSize / blockSize
	return &Cache{
		numLines:  numLines,
		blockSize: blockSize,
		tags:      make([]int, numLines),
		valid:     make([]bool, numLines),
	}
}

func (c *Cache) SimulateAccess(address int) (hit bool) {
	// blockOffset := address % c.blockSize
	index := (address / c.blockSize) % c.numLines
	tag := address / (c.blockSize * c.numLines)

	if c.tags[index] == tag && c.valid[index] {
		return true // hit
	}

	// Miss: update cache line
	c.tags[index] = tag
	c.valid[index] = true
	return false
}
