package hashtable

// 哈希表
func MinNumberDisappeared1(nums []int) int {
	m := map[int]bool{}
	for _, num := range nums {
		if num > 0 {
			m[num] = true
		}
	}
	for i := 1; i <= len(nums); i++ {
		if !m[i] {
			return i
		}
	}
	return len(nums) + 1
}

// 原地哈希
func MinNumberDisappeared2(nums []int) int {
	for i, num := range nums {
		if num <= 0 {
			nums[i] = len(nums) + 1
		}
	}
	for _, num := range nums {
		if Abs(num) <= len(nums) {
			nums[Abs(num)-1] = -nums[Abs(num)-1]
		}
	}
	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return len(nums) + 1
}
