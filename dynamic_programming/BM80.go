package dynamicprogramming

// 遍历，遍历的过程中保存已经遍历过的元素的最小值，当前的元素减去它就可以得到在这天卖出股票得到的最大收益。
// 在每天卖出股票能够获得的最大收益取去最大值即为所求
func MaxProfitBM80_1(prices []int) int {
	low, maxProfit := 10001, 0
	for _, price := range prices {
		if price > low {
			maxProfit = Max(maxProfit, price-low)
		} else {
			low = Min(low, price)
		}
	}
	return maxProfit
}

// 动态规划
// 设 dp[i] 表示在第 i 天卖出股票获得的最大收益
// dp[i] = price[i-1] - min(price[i-2]-dp[i-1], price[i-2])
// dp[0] = 0, dp[1] = 0
func MaxProfitBM80_2(prices []int) int {
	a, maxProfit := 0, 0
	for i := 2; i <= len(prices); i++ {
		a = prices[i-1] - Min(prices[i-2]-a, prices[i-2])
		maxProfit = Max(maxProfit, a)
	}
	return maxProfit
}
