package binarysearchsort

// 双指针，在遍历的过程中计算 version1 和 version2 的修订号，当遇到 '.' 的时候就比较修订号大小。
// 当其中一个指针遍历到结尾的情况下有两种情况（假设遍历到结尾的是version1）：
//
// 1. version2 在相同版本段的修订号更长
//
// 2. version2 中的 '.' 数量多于 version1
//
// 对于上述两种情况下，都需要继续计算 version2 中相同版本段的修订号码，
// 如果是第 2 种情况，在修订号相同的情况下还需要判断 version2 中 '.' 的后面是否有 1，如果是则 version2 更大。
func Compare(version1 string, version2 string) int {
	v1, v2 := 0, 0 // 修订号
	i, j := 0, 0   // 下标指针

	for i < len(version1) && j < len(version2) {
		if version1[i] == '.' && version2[j] == '.' {
			if v1 > v2 {
				return 1
			} else if v1 < v2 {
				return -1
			}
			v1, v2 = 0, 0
			i++
			j++
		} else if version1[i] == '.' {
			v2 = v2*10 + int(version2[j]-'0')
			j++
		} else if version2[j] == '.' {
			v1 = v1*10 + int(version1[i]-'0')
			i++
		} else {
			v1 = v1*10 + int(version1[i]-'0')
			v2 = v2*10 + int(version2[j]-'0')
			i++
			j++
		}
	}
	for i < len(version1) && version1[i] != '.' {
		v1 = v1*10 + int(version1[i]-'0')
		i++
	}
	for j < len(version2) && version2[j] != '.' {
		v2 = v2*10 + int(version2[j]-'0')
		j++
	}
	if v1 > v2 {
		return 1
	} else if v1 < v2 {
		return -1
	}
	for i < len(version1) {
		if version1[i] == '1' {
			return 1
		}
		i++
	}
	for j < len(version2) {
		if version2[j] == '1' {
			return -1
		}
		j++
	}

	return 0
}
