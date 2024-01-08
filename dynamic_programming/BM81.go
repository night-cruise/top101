package dynamicprogramming

// 动态规划 (状态转移)
// 设 own[i] 表示第 i 天拥有股票的最大收益
// 设 los[i] 表示第 i 天不拥有股票的最大收益
// own[i] = max(own[i-1], los[i-1] - prices[i-1])
// los[i] = max(los[i-1], own[i-1] + prices[i-1])
// own[1] = -prices[0], los[1] = 0
func MaxProfitBM81(prices []int) int {
	a, b := -prices[0], 0
	for i := 2; i <= len(prices); i++ {
		a, b = Max(a, b-prices[i-1]), Max(b, a+prices[i-1])
	}
	return b
}
