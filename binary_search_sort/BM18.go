package binarysearchsort

// 线性搜索
func Find1(target int, array [][]int) bool {
	if len(array) == 0 || len(array[0]) == 0 {
		return false
	}
	i, j := 0, len(array[0])-1

	for i < len(array) && j >= 0 {
		if target == array[i][j] {
			return true
		} else if target > array[i][j] {
			i++
		} else {
			j--
		}
	}

	return false
}

// 二分查找
func Find2(target int, array [][]int) bool {
	if len(array) == 0 || len(array[0]) == 0 {
		return false
	}

	for i := 0; i < len(array); i++ {
		left, right := 0, len(array[i])
		for left < right {
			mid := left + (right-left)/2
			if target > array[i][mid] {
				left = mid + 1
			} else {
				right = mid
			}
			if left != len(array[i]) && target == array[i][left] {
				return true
			}
		}
	}

	return false
}
