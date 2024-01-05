package hashtable

// 哈希表
func FindNumsAppearOnce1(nums []int) []int {
	m := map[int]int{}

	for _, num := range nums {
		m[num]++
	}

	res := []int{}
	for num, count := range m {
		if count == 1 {
			res = append(res, num)
		}
	}

	if res[0] > res[1] {
		res[0], res[1] = res[1], res[0]
	}
	return res
}

// 异或，利用相同的数字的异或结果为 0
func FindNumsAppearOnce2(nums []int) []int {
	xorAll := 0
	for _, num := range nums {
		xorAll ^= num
	}
	seq := 1
	for seq&xorAll == 0 {
		seq <<= 1
	}
	a, b := 0, 0
	for _, num := range nums {
		if num&seq > 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	if a > b {
		a, b = b, a
	}
	return []int{a, b}
}
