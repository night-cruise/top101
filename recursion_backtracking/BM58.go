package recursionbacktracking

import "sort"

// 排序 + 回溯
//
//  Permutation(str) = str[i] + Permutation(str - str[i])
func PermutationBM58(str string) []string {
	s := []byte(str)
	sort.Slice(s, func(i, j int) bool { return s[i] <= s[j] })
	path, m, ans := []byte{}, make([]bool, len(s)), []string{}
	var dfsHelper func()

	dfsHelper = func() {
		if len(path) == len(s) {
			ans = append(ans, string(path))
			return
		}
		for i := 0; i < len(s); i++ {
			if m[i] {
				continue
			}
			m[i] = true
			path = append(path, s[i])
			dfsHelper()
			m[i] = false
			path = path[:len(path)-1]

			for i < len(s)-1 && s[i] == s[i+1] {
				i++
			}
		}
	}
	dfsHelper()

	return ans
}
