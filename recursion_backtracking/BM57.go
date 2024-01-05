package recursionbacktracking

// DFS
func Solve(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '2' || grid[i][j] == '0' {
				continue
			}
			DFSHelper(grid, i, j)
			ans++
		}
	}
	return ans
}

func DFSHelper(grid [][]byte, x, y int) {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
		return
	}
	if grid[x][y] != '1' {
		return
	}
	grid[x][y] = '2'
	DFSHelper(grid, x-1, y)
	DFSHelper(grid, x+1, y)
	DFSHelper(grid, x, y-1)
	DFSHelper(grid, x, y+1)
}
