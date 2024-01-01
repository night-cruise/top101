package binarysearchsort

// 二分查找
//
// 一般情况：3, 4, 5, 1, 2，如果 nums[mid] > nums[right]，那么最小元素一定在nums(mid+1..right)中，否则最小元素在nums(..mid)中。
//
// 特殊情况：2, 2, 2, 1, 2，最小元素一定在 nums(left+1..right-1)中。
func MinNumberInRotateArray(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		}
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
