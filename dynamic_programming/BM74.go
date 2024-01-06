package dynamicprogramming

import (
	"strconv"
	"strings"
)

// 回溯
// RestoreIpAddresses(s) = IpNew(s[0:i]) + RestoreIpAddresses(s[i:])
func RestoreIpAddresses(s string) []string {
	path, ans := []string{}, []string{}
	var backtracking func(i int)

	backtracking = func(i int) {
		if i == len(s) && len(path) == 4 {
			ans = append(ans, strings.Join(path, "."))
			return
		}
		if i+1 > len(s) {
			return
		}
		path = append(path, string(s[i:i+1]))
		backtracking(i + 1)
		path = path[:len(path)-1]

		if i+2 > len(s) || s[i] == '0' {
			return
		}
		path = append(path, string(s[i:i+2]))
		backtracking(i + 2)
		path = path[:len(path)-1]

		if i+3 > len(s) || s[i] == '0' {
			return
		}
		if val, _ := strconv.Atoi(string(s[i : i+3])); val <= 255 {
			path = append(path, string(s[i:i+3]))
			backtracking(i + 3)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)

	return ans
}
