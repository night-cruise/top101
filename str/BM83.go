package str

import "strings"

// 双指针
func Trans(s string, n int) string {
	strList := strings.Split(s, " ")
	i, j := 0, len(strList)-1
	for i < j {
		s1, s2 := []byte(strList[i]), []byte(strList[j])
		for k := 0; k < len(s1); k++ {
			if s1[k] >= 'a' && s1[k] <= 'z' {
				s1[k] = 'A' + s1[k] - 'a'
			} else {
				s1[k] = 'a' + s1[k] - 'A'
			}
		}
		for k := 0; k < len(s2); k++ {
			if s2[k] >= 'a' && s2[k] <= 'z' {
				s2[k] = 'A' + s2[k] - 'a'
			} else {
				s2[k] = 'a' + s2[k] - 'A'
			}
		}
		strList[i], strList[j] = string(s2), string(s1)
		i++
		j--
	}
	if i == j {
		str := []byte(strList[i])
		for k := 0; k < len(str); k++ {
			if str[k] >= 'a' && str[k] <= 'z' {
				str[k] = 'A' + str[k] - 'a'
			} else {
				str[k] = 'a' + str[k] - 'A'
			}
		}
		strList[i] = string(str)
	}
	return strings.Join(strList, " ")
}
