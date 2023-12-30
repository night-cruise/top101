package binarysearchsort

// 这道题给的不严谨，题目中没有限定所有的元素一定是相等的

// 遍历
func FindPeakElement1(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[len(nums)-1] > nums[len(nums)-2] {
		return len(nums) - 1
	}
	return -1
}

// 遍历，最大元素的下标即为所求
func FindPeakElement2(nums []int) int {
	res := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[res] {
			res = i
		}
	}

	return res
}

// 二分查找
func FindPeakElement3(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid+1] > nums[mid] { // 山峰一定在右边
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
