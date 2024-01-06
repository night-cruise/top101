package dynamicprogramming

// 暴力解，三重循环
func GetLongestPalindrome1(A string) int {
	maxLen := 1
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			maxLen = Max(maxLen, GPHelper(A, i, j))
		}
	}
	return maxLen
}
func GPHelper(A string, i, j int) int {
	length := 0
	for i < j {
		if A[i] != A[j] {
			break
		}
		length += 2
		i++
		j--
	}
	if i < j {
		return 0
	}
	if i == j {
		length++
	}
	return length
}

// 动态规划
// dp[i][j] 表示字符串 A[i:j+1] 的最长回文子串的长度
// dp[i][j] = if dp[i+1][j-1] == j-i-1 && A[i] == A[j] { dp[i+1][j-1] + 2 } else { Max(dp[i+1][[j], dp[i][j-1]) }
// dp[0][0] = 1, dp[len(A)-1][len(A)-1] = 1
func GetLongestPalindrome2(A string) int {
	dp := make([][]int, len(A))
	for i := range dp {
		dp[i] = make([]int, len(A))
	}
	dp[0][0], dp[len(A)-1][len(A)-1] = 1, 1
	for i := len(A) - 2; i >= 0; i-- {
		for j := i; j < len(A); j++ {
			if j == i {
				dp[i][j] = 1
				continue
			}
			if A[i] == A[j] && dp[i+1][j-1] == j-i-1 {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = Max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(A)-1]
}

// 中心扩散法，遍历字符串 A，然后以当前位置 (i, i) 或者 (i,i+1) 为中心向两端扩散，计算回文串的长度
func GetLongestPalindrome3(A string) int {
	getPalindrome := func(A string, i, j int) int {
		length := 0
		for i >= 0 && j < len(A) && A[i] == A[j] {
			if i == j {
				length++
			} else {
				length += 2
			}
			i--
			j++
		}
		return length
	}
	maxLen := 0
	for i := 0; i < len(A); i++ {
		maxLen = Max(maxLen, getPalindrome(A, i, i))
		maxLen = Max(maxLen, getPalindrome(A, i, i+1))
	}
	return maxLen
}
