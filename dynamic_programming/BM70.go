package dynamicprogramming

// 动态规划
// dp[n][aim] 表示在货币数组 coins[0:n] 中找零 aim 需要的最少的货币数量
// dp[n][aim] = min(dp[n][aim-coins[n-1]] + 1, dp[n-1][aim])
// dp[i][0] = 0, dp[0][j] = 5001
func MinMoney(arr []int, aim int) int {
	dp := make([][]int, len(arr)+1)
	for i := range dp {
		dp[i] = make([]int, aim+1)
	}
	for j := 1; j <= aim; j++ {
		dp[0][j] = 5001
	}
	for i := 1; i <= len(arr); i++ {
		for j := 1; j <= aim; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= arr[i-1] {
				dp[i][j] = Min(dp[i][j], dp[i][j-arr[i-1]]+1)
			}
		}
	}
	if dp[len(arr)][aim] > 5000 {
		return -1
	}
	return dp[len(arr)][aim]
}
