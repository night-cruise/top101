package binarysearchsort

// 两层遍历
func InversePairs1(nums []int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				sum++
			}
		}
	}

	return sum % 1000000007
}

// 归并统计，使用归并排序的思想，当处于合并阶段的时候可以统计逆序对的数量
func InversePair2(nums []int) int {
	tmp := make([]int, len(nums))
	return inverseParisHelper(nums, tmp) % 1000000007
}

func inverseParisHelper(nums, tmp []int) int {
	if len(nums) <= 1 {
		return 0
	}
	mid := len(nums) / 2
	leftInverseParis, rightInversePairs := inverseParisHelper(nums[:mid], tmp), inverseParisHelper(nums[mid:], tmp)

	i, j, k := 0, mid, 0
	curInversePairs := 0

	for i < mid && j < len(nums) {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			curInversePairs += mid - i
			tmp[k] = nums[j]
			j++
		}
		k++
	}
	for i < mid {
		tmp[k] = nums[i]
		i++
		k++
	}
	for j < len(nums) {
		tmp[k] = nums[j]
		j++
		k++
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
	return curInversePairs + leftInverseParis + rightInversePairs
}
