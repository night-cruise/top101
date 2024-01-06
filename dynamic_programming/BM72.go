package dynamicprogramming

// 动态规划
// dp[n] 表示数组 array[:n] 以 array[n-1] 结尾的子数组的和的最大值
// dp[n] = if dp[n-1] > 0 { dp[n-1] + array[n-1] } else { array[n-1] }
// dp[0] = -101
func FindGreatestSumOfSubArray1(array []int) int {
	dp := make([]int, len(array)+1)
	dp[0] = -101
	maxSum := -101

	for i := 1; i <= len(array); i++ {
		dp[i] = Max(dp[i-1]+array[i-1], array[i-1])
		maxSum = Max(maxSum, dp[i])
	}

	return maxSum
}

// 动态规划（空间优化）
func FindGreatestSumOfSubArray2(array []int) int {
	cur, maxSum := -101, -101

	for i := 1; i <= len(array); i++ {
		cur = Max(cur+array[i-1], array[i-1])
		maxSum = Max(maxSum, cur)
	}

	return maxSum
}
