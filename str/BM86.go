package str

import "strconv"

// 模拟 + 双指针
func SolveBM86(s string, t string) string {
	ans := ""
	carry := 0
	i, j := len(s)-1, len(t)-1
	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(s[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(t[j] - '0')
			j--
		}
		ans = strconv.Itoa(sum%10) + ans
		carry = sum / 10
	}
	if carry == 1 {
		ans = "1" + ans
	}
	return ans
}
