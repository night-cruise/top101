package dynamicprogramming

// 暴力解，三重循环
func LCSBM66_1(str1 string, str2 string) string {
	ans := ""
	for i := 0; i < len(str1); i++ {
		for j := 0; j < len(str2); j++ {
			if str1[i] != str2[j] {
				continue
			}
			k, m := i+1, j+1
			for k < len(str1) && m < len(str2) && str1[k] == str2[m] {
				k++
				m++
			}
			if k-i <= len(ans) {
				continue
			}
			ans = string(str1[i:k])
		}
	}
	return ans
}

// 动态规划
// dp[i][j] 为字符串 s1[0:i] 和字符串 s2[0:j] 以字符 s1[i-1] 和字符 s2[j-1] 结尾的最长公共子串的长度
// dp[i][j] = if s1[i-1] == s2[j-1] { dp[i-1][j-1] + 1 } else { 0 }
// dp[0][j] = dp[i][0] = 0
func LCSBM66_2(str1 string, str2 string) string {
	dp := make([][]int, len(str1)+1)
	for i := range dp {
		dp[i] = make([]int, len(str2)+1)
	}
	ans := ""
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] != str2[j-1] {
				continue
			}
			dp[i][j] = dp[i-1][j-1] + 1
			if dp[i][j] > len(ans) {
				ans = string(str1[i-dp[i][j] : i])
			}
		}
	}
	return ans
}
