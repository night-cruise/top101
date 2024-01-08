package dynamicprogramming

// 动态规划，分两种情况考虑：（1）不偷第一间房子（2）不偷最后一间房子
// 设dp[i]表示经过第i间房子时能够盗取的最大金额
// dp[i] = max(dp[i-1], d[i-2] + nums[i-1])
// dp[0] = 0, dp[1] = nums[0]
func RobBM79(nums []int) int {
	dynamic := func(nums []int) int {
		a, b := 0, nums[0]
		for i := 2; i <= len(nums); i++ {
			a, b = b, Max(b, a+nums[i-1])
		}
		return b
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return Max(dynamic(nums[:len(nums)-1]), dynamic(nums[1:]))
}
