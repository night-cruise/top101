package dynamicprogramming

// 动态规划
// dp[n] 表示数组 arr[:n] 以 arr[n-1] 结尾的最长严格上升子列的长度
// dp[n] = max(1, dp[i] + 1), arr[n-1] > arrp[i-1]
// dp[0] = 0
func LIS(arr []int) int {
	dp := make([]int, len(arr)+1)
	maxLen := 0

	for i := 1; i <= len(arr); i++ {
		dp[i] = 1
		for j := i - 1; j >= 1; j-- {
			if arr[i-1] > arr[j-1] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}
		maxLen = Max(maxLen, dp[i])
	}

	return maxLen
}
