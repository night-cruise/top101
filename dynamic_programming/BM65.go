package dynamicprogramming

// 动态规划
// dp[i][j] 为 s1[0:i] 和 s2[0:j] 的最长公共子序列字符串
// dp[i][j] = if s[i-1] == s[j-1] { dp[i-1][j-1] + 1 } else { max(dp[i-1][j], dp[i][j-1]) }
// dp[0][j] = 0, dp[i][0] = 0
func LCSBM65(s1 string, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return "-1"
	}
	dp := make([][]string, len(s1)+1)
	for i := range dp {
		dp[i] = make([]string, len(s2)+1)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + string(s1[i-1])
			} else {
				if len(dp[i-1][j]) > len(dp[i][j-1]) {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	if dp[len(s1)][len(s2)] == "" {
		return "-1"
	}
	return dp[len(s1)][len(s2)]
}
