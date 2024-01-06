package dynamicprogramming

// 动态规划
// dp[i][j] 表示将字符串 str1[i] 转化为 str2[j] 的最小操作次数
// dp[i][j] = dp[i-1][j-1], str1[i-1] == str2[j-1]
//          = dp[i-1][j-1] + 1, str1[i-1] -> str2[j-1]
//			= dp[i-1][j] + 1, del str1[i-1]
//          = dp[i][j-1] + 1, add str2[j-1]
// 后面三种情况取最小值
// dp[0][j] = j, dp[i][0] = i
func EditDistance(str1 string, str2 string) int {
	dp := make([][]int, len(str1)+1)
	for i := range dp {
		dp[i] = make([]int, len(str2)+1)
	}
	for i := 1; i <= len(str1); i++ {
		dp[i][0] = i
	}
	for j := 1; j <= len(str2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = Min(Min(dp[i-1][j-1], dp[i-1][j]), dp[i][j-1]) + 1
			}
		}
	}
	return dp[len(str1)][len(str2)]
}
