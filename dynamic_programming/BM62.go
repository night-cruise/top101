package dynamicprogramming

// dp[n] = dp[n-1] + dp[n-2], n > 2
// dp[1] = dp[2] = 1
func Fibonacci(n int) int {
	if n <= 2 {
		return 1
	}
	a, b := 1, 1
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}
