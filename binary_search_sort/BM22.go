package binarysearchsort

// 双指针
func Compare(version1 string, version2 string) int {
	len1, len2 := 0, 0

	for len1 < len(version1) || len2 < len(version2) {
		num1, num2 := 0, 0

		for len1 < len(version1) && version1[len1] != '.' {
			num1 = num1*10 + int(version1[len1]-'0')
			len1++
		}
		for len2 < len(version2) && version2[len2] != '.' {
			num2 = num2*10 + int(version2[len2]-'0')
			len2++
		}
		if num1 > num2 {
			return 1
		}
		if num2 > num1 {
			return -1
		}
		len1++
		len2++
		num1, num2 = 0, 0
	}

	return 0
}
