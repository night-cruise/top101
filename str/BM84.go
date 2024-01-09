package str

// 遍历
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	length := 0
	for {
		flag := false
		for j := 0; j < len(strs); j++ {
			if length >= len(strs[j]) || j >= 1 && strs[j][length] != strs[j-1][length] {
				flag = true
				break
			}
		}
		if flag {
			break
		}
		length++
	}
	return strs[0][:length]
}
