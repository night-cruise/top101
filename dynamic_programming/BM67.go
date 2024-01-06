package dynamicprogramming

// 动态规划
// dp[m][n] 为从 (0, 0) 到 (m, n) 的不同路径的数量。
// dp[m][n] = dp[m-1][n] + dp[m][n-1]
// dp[0][n] = 1, dp[m][0] = 1
func UniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 组合数学
// 从 (0, 0) 移动到 (m-1, n-1) 需要向右移动 n-1 步，向下移动 m-1 步，组合起来一共存在 C(m+n-2,m-1) 个不同的不经
// c(m+n-2,m-1) = (m+n-2)! / ((m-1)! * (n-1)!) = (m+n-2)*....*n / (m-1)!
func UniquePaths2(m int, n int) int {
	x, y := 1, 1

	for i := 2; i <= m+n-2; i++ {
		if i <= m-1 && i >= n {
			continue
		}
		if i <= m-1 {
			y *= i
		}
		if i >= n {
			x *= i
		}
	}

	return x / y
}
