package hashtable

import "sort"

// 哈希表
func MoreThanHalfNum_Solution1(numbers []int) int {
	m := map[int]int{}

	for _, num := range numbers {
		m[num]++
	}

	for num, cnt := range m {
		if cnt > len(numbers)/2 {
			return num
		}
	}

	return -1
}

// 排序
func MoreThanHalfNum_Solution2(numbers []int) int {
	sort.Slice(numbers, func(i, j int) bool { return numbers[i] < numbers[j] })

	return numbers[len(numbers)/2]
}

// 摩尔投票算法
func MoreThanHalfNum_Solution3(numbers []int) int {
	candidate, count := numbers[0], 1

	for i := 1; i < len(numbers); i++ {
		if numbers[i] == candidate {
			count++
		} else {
			count--
			if count == 0 {
				candidate = numbers[i]
				count = 1
			}
		}
	}

	return candidate
}
