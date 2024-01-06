package dynamicprogramming

// 动态规划
// dp[i][j] 为从 (0, 0) 到 （i, j） 的最小路径和
// dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
// dp[0][0] = matrix[0][0], dp[0][j] = dp[0][j-1] + matrix[0][j], dp[i][0] = dp[i-1][0] + matrix[i][0]
func MinPathSum(matrix [][]int) int {
	dp := make([][]int, len(matrix))
	for i := range dp {
		dp[i] = make([]int, len(matrix[0]))
	}
	dp[0][0] = matrix[0][0]
	for i := 1; i < len(matrix); i++ {
		dp[i][0] = dp[i-1][0] + matrix[i][0]
	}
	for j := 1; j < len(matrix[0]); j++ {
		dp[0][j] = dp[0][j-1] + matrix[0][j]
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			dp[i][j] = Min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
		}
	}
	return dp[len(matrix)-1][len(matrix[0])-1]
}
