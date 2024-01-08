package dynamicprogramming

// https://blog.nowcoder.net/n/9e4b3953abb84149b04c95bf89d7772e

// 暴力解，遍历字符串 s，计算以 s[i] 开头的最长的格式正确的括号子串的长度
func LongestValidParentheses1(s string) int {
	maxLen := 0
	for i := 0; i < len(s); i++ {
		leftBrackets, rightBrackets := 0, 0
		for j := i; j < len(s); j++ {
			if s[j] == '(' {
				leftBrackets++
			} else {
				rightBrackets++
			}
			if leftBrackets == rightBrackets {
				maxLen = Max(2*leftBrackets, maxLen)
			} else if rightBrackets > leftBrackets {
				break
			}
		}
	}
	return maxLen
}

// 动态规划
// dp[i] 表示字符串 s[:i] 以字符 s[i-1] 结尾的最长的格式正确的括号子串的长度
// dp[i] = 0, s[i-1] == '('
//	     = dp[i-2] + 2, s[i-1] == ')' && s[i-2] == '('
// 		 = dp[i-1] + dp[i - 2 - dp[i-1]] + 2, s[i-1] == ')' && s[i - 2 - dp[i-1]] == '('
//       = 0, s[i-1] == ')' && s[i - 2 - dp[i-1]] != '('
// dp[0] = 0, dp[1] = 0
func LongestValidParentheses2(s string) int {
	dp := make([]int, len(s)+1)
	maxLen := 0
	for i := 2; i <= len(s); i++ {
		if s[i-1] == '(' {
			continue
		}
		if s[i-2] == '(' {
			dp[i] = dp[i-2] + 2
		} else if i-2 < dp[i-1] || s[i-2-dp[i-1]] != '(' {
			continue
		} else {
			dp[i] = dp[i-1] + dp[i-2-dp[i-1]] + 2
		}
		maxLen = Max(maxLen, dp[i])
	}
	return maxLen
}

// 栈，栈底保存当前已经遍历过的元素中[最后一个没有匹配右括号的下标]
func LongestValidParentheses3(s string) int {
	stack, maxLen := []int{-1}, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		if len(stack) >= 2 {
			stack = stack[:len(stack)-1]
			maxLen = Max(maxLen, i-stack[len(stack)-1])
		} else {
			stack[0] = i
		}
	}
	return maxLen
}

// 正反遍历 + 双计数器
func LongestValidParentheses4(s string) int {
	maxLen := 0

	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLen = Max(maxLen, 2*left)
		} else if right > left {
			left, right = 0, 0
		}
	}

	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLen = Max(maxLen, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}

	return maxLen
}
