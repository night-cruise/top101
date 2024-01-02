package binarysearchsort

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
		if nums[mid+1] > nums[mid] {
			left = mid + 1 // mid 的右边一定存在山峰
		} else {
			right = mid // mid 的左边一定存在山峰
		}
	}

	return left
}
