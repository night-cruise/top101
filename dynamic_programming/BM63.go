package dynamicprogramming

// 动态规划
// dp[n] 表示跳上 n 阶台阶的跳法数
// dp[n] = dp[n-2] + dp[n-1]
// dp[1] = 1, dp[2] = 2
func JumpFloor(number int) int {
	a, b := 1, 1
	for i := 2; i <= number; i++ {
		a, b = b, a+b
	}
	return b
}
