package binarysearchsort

func Search(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	if left == len(nums) || nums[left] != target {
		return -1
	}
	return left
}
