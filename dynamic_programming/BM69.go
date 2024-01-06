package dynamicprogramming

// 动态规划
// dp[n] 表示字符串 nums[:n] 的可能的编码数量
// dp[n] = if nums[n-1] != '0' { dp[n-1] } + if nums[n-2] != '0' && int(nums[n-2:n] <= 26) { dp[n-2] }
// dp[0] = 1, dp[1] = if nums[0] == '0' { 0 } else { 1 }
func SolveBM69(nums string) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 1
	if nums[0] != '0' {
		dp[1] = 1
	}

	for i := 2; i <= len(nums); i++ {
		if nums[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if nums[i-2] != '0' && (10*(nums[i-2]-'0')+nums[i-1]-'0') <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(nums)]
}
