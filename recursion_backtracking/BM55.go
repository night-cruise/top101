package recursionbacktracking

// 回溯
//
// Permute(num) = num[i] + Permute(num - num[i])
func PermuteBM55(num []int) [][]int {
	m, cur, ans := make([]bool, len(num)), []int{}, [][]int{}
	var pHelper func()

	pHelper = func() {
		if len(cur) == len(num) {
			ans = append(ans, append([]int{}, cur...))
			return
		}
		for i := range num {
			if m[i] {
				continue
			}
			m[i] = true
			cur = append(cur, num[i])
			pHelper()
			m[i] = false
			cur = cur[:len(cur)-1]
		}
	}
	pHelper()

	return ans
}
