package recursionbacktracking

// 暴力搜索，搜索以矩阵中的每个位置为起始位置的最长递增路径的长度，然后取最大值。
func Solve1(matrix [][]int) int {
	maxLength := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			maxLength = Max(maxLength, SHelper1(matrix, i, j, -1))
		}
	}
	return maxLength
}

func SHelper1(matrix [][]int, x, y int, pre int) int {
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) {
		return 0
	}
	if matrix[x][y] <= pre {
		return 0
	}
	length := 0
	length = Max(length, 1+SHelper1(matrix, x-1, y, matrix[x][y]))
	length = Max(length, 1+SHelper1(matrix, x+1, y, matrix[x][y]))
	length = Max(length, 1+SHelper1(matrix, x, y-1, matrix[x][y]))
	length = Max(length, 1+SHelper1(matrix, x, y+1, matrix[x][y]))
	return length
}

// 记忆化搜索
func Solve2(matrix [][]int) int {
	visited, maxLength := make([][]int, len(matrix)), 0
	for i := range matrix {
		visited[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			maxLength = Max(maxLength, SHelper2(matrix, i, j, -1, visited))
		}
	}
	return maxLength
}

func SHelper2(matrix [][]int, x, y int, pre int, visited [][]int) int {
	if x < 0 || x >= len(matrix) || y < 0 || y >= len(matrix[0]) {
		return 0
	}
	if matrix[x][y] <= pre {
		return 0
	}
	if visited[x][y] != 0 {
		return visited[x][y]
	}
	length := 0
	length = Max(length, 1+SHelper2(matrix, x-1, y, matrix[x][y], visited))
	length = Max(length, 1+SHelper2(matrix, x+1, y, matrix[x][y], visited))
	length = Max(length, 1+SHelper2(matrix, x, y-1, matrix[x][y], visited))
	length = Max(length, 1+SHelper2(matrix, x, y+1, matrix[x][y], visited))
	visited[x][y] = length
	return length
}
