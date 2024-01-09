package dynamicprogramming

// 动态规划
// 设 own1[i] 表示第 1 笔交易第 i+1 天拥有股票的最大收益
// 设 los1[i] 表示第 1 笔交易第 i+1 天不拥有股票的最大收益
// 设 own2[i] 表示第 2 笔交易第 i+1 天拥有股票的最大收益
// 设 los2[i] 表示第 2 笔交易第 i+1 天不拥有股票的最大收益
// own1[i] = max(own1[i-1], -prices[i])
// los1[i] = max(los1[i-1], own1[i-1] + prices[i])
// own2[i] = max(own2[i-1], los1[i-1] - prices[i])
// los2[i] = max(los2[i-1], own2[i-1] + prices[i])
// own1[0] = -prices[0], los1[0] = 0, own2[0] = -prices[0], los2[0] = 0
func MaxProfitBM82_1(prices []int) int {
	own1, los1, own2, los2 := make([]int, len(prices)), make([]int, len(prices)), make([]int, len(prices)), make([]int, len(prices))
	own1[0], own2[0] = -prices[0], -prices[0]
	for i := 1; i < len(prices); i++ {
		own1[i] = Max(own1[i-1], -prices[i])
		los1[i] = Max(los1[i-1], own1[i-1]+prices[i])
		own2[i] = Max(own2[i-1], los1[i-1]-prices[i])
		los2[i] = Max(los2[i-1], own2[i-1]+prices[i])
	}
	return los2[len(prices)-1]
}

// 动态规划（空间压缩）
func MaxProfitBM82_2(prices []int) int {
	own1, los1, own2, los2 := -prices[0], 0, -prices[0], 0
	for i := 1; i < len(prices); i++ {
		own1, los1, own2, los2 = Max(own1, -prices[i]), Max(los1, own1+prices[i]), Max(own2, los1-prices[i]), Max(los2, own2+prices[i])
	}
	return los2
}
