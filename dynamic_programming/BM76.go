package dynamicprogramming

// 动态规划
// dp[i][j] 表示字符串 str[:i] 与模式 pattern[:j] 是否匹配
// dp[i][j] = dp[i-1][j-1], str[i-1] == pattern[j-1] || pattern[j-1] == '.'
// 		    = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j-1] || dp[i-1][j], pattern[j-1] == '*' &&  (pattern[j-2] == '.' || str[i-1] == pattern[j-2])
//          = dp[i][j-2] || dp[i][j-1], pattern[j-1] == '*' && pattern[j-2] == '*'
//          = dp[i][j-2], pattern[j-1] == '*' && pattern[j-2] != '*' && pattern[j-2] != '.' && str[i-1] != pattern[j-2]
// dp[0][0] = true, dp[0][j] = if pattern[j-1] == '*' { dp[0][j-1] || dp[0][j-2] } else { false }, dp[i][0] = false
func Match(str string, pattern string) bool {
	dp := make([][]bool, len(str)+1)
	for i := range dp {
		dp[i] = make([]bool, len(pattern)+1)
	}
	dp[0][0] = true
	for j := 1; j <= len(pattern); j++ {
		if pattern[j-1] == '*' {
			dp[0][j] = dp[0][j-2] || dp[0][j-1]
		}
	}
	for i := 1; i <= len(str); i++ {
		for j := 1; j <= len(pattern); j++ {
			if str[i-1] == pattern[j-1] || pattern[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if pattern[j-1] == '*' {
				if pattern[j-2] == '.' || str[i-1] == pattern[j-2] {
					dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j-1] || dp[i-1][j]
				} else if pattern[j-2] == '*' {
					dp[i][j] = dp[i][j-2] || dp[i][j-1]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[len(str)][len(pattern)]
}
