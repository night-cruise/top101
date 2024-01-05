package hashtable

import "sort"

// 暴力解法，排序 + 三重循环
func ThreeSum1(num []int) [][]int {
	sort.Ints(num)
	ans := [][]int{}

	for i := 0; i < len(num); i++ {
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		for j := i + 1; j < len(num); j++ {
			if j > i+1 && num[j] == num[j-1] {
				continue
			}
			for k := j + 1; k < len(num); k++ {
				if k > j+1 && num[k] == num[k-1] {
					continue
				}
				if num[i]+num[j]+num[k] == 0 {
					ans = append(ans, []int{num[i], num[j], num[k]})
					break
				}
			}
		}
	}

	return ans
}

// 优化，排序 + 双指针
func ThreeSum2(num []int) [][]int {
	sort.Ints(num)
	ans := [][]int{}

	for i := 0; i < len(num); i++ {
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		j, k := i+1, len(num)-1
		for j < k {
			if num[j]+num[k] == -num[i] {
				ans = append(ans, []int{num[i], num[j], num[k]})
				j++
				k--
			} else if num[j]+num[k] < -num[i] {
				j++
			} else {
				k--
			}
			for j > i+1 && j < len(num) && num[j] == num[j-1] {
				j++
			}
			for k < len(num)-1 && k >= 0 && num[k] == num[k+1] {
				k--
			}
		}
	}

	return ans
}
