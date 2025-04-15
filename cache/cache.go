package cache

// traversal to test spacial locality

const (
	Rows = 1024
	Cols = 1024
)

func RowWiseTraversal(matrix [][]int) int {
	sum := 0
	for i := 0; i < Rows; i++ {
		for j := 0; j < Cols; j++ {
			sum += matrix[i][j] //  Good spatial locality
		}
	}

	return sum
}

func ColumTraversal(matrix [][]int) int {
	sum := 0

	for j := 0; j < Cols; j++ {
		for i := 0; i < Rows; i++ {
			sum += matrix[i][j]
		}
	}

	return sum
}
