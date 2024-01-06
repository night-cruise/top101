package dynamicprogramming

// 动态规划
// dp[n] 表示爬 n 阶楼梯需要支付的最低费用。
// dp[n] = min(dp[n-2] + cost[n-2], dp[n-1] + cost[n-1])
// dp[0] = 0, dp[1] = 0
func MinCostClimbingStairs(cost []int) int {
	if len(cost) <= 1 {
		return 0
	}
	a, b := 0, 0
	for i := 2; i <= len(cost); i++ {
		a, b = b, Min(a+cost[i-2], b+cost[i-1])
	}
	return b
}
