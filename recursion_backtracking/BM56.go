package recursionbacktracking

import "sort"

// 排序 + 回溯
//
// PermuteUnique(num) = num[i] + PermuteUnique(num - num[i])
func PermuteUnique(num []int) [][]int {
	sort.Ints(num)
	m, path, ans := make([]bool, len(num)), []int{}, [][]int{}
	var pHelper func()

	pHelper = func() {
		if len(path) == len(num) {
			ans = append(ans, append([]int{}, path...))
			return
		}
		for i := 0; i < len(num); i++ {
			if m[i] {
				continue
			}
			m[i] = true
			path = append(path, num[i])
			pHelper()
			m[i] = false
			path = path[:len(path)-1]

			for i < len(num)-1 && num[i] == num[i+1] {
				i++
			}
		}
	}
	pHelper()

	return ans
}
