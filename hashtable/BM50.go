package hashtable

// 哈希表
func TwoSum(numbers []int, target int) []int {
	m := map[int]int{}
	twoNums := make([]int, 2)

	for i, num := range numbers {
		if j, ok := m[target-num]; ok {
			twoNums[0], twoNums[1] = j+1, i+1
			break
		}
		m[num] = i
	}

	return twoNums
}
