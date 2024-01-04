package heapstackqueue

import (
	"container/heap"
	"sort"
)

// 排序
func GetLeastNumbers_Solutio1(input []int, k int) []int {
	if len(input) == 0 || k == 0 {
		return []int{}
	}
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })
	return input[:k]
}

// 大顶堆，时间复杂度 O(NlogK)
func GetLeastNumbers_Solutio2(input []int, k int) []int {
	if len(input) == 0 || k == 0 {
		return []int{}
	}
	h := &MaxHeap{}

	for _, item := range input {
		heap.Push(h, item)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return []int(*h)
}

// 快速选择，时间复杂度 O(N)
func GetLeastNumbers_Solutio3(input []int, k int) []int {
	if len(input) == 0 || k == 0 {
		return []int{}
	}
	quickSelect(input, k)

	return input[:k]
}

func quickSelect(input []int, k int) {
	if len(input) <= 1 {
		return
	}
	i, j, privot := -1, len(input), input[0]

	for i < j {
		for {
			i++
			if input[i] >= privot {
				break
			}
		}
		for {
			j--
			if input[j] <= privot {
				break
			}
		}
		if i < j {
			input[i], input[j] = input[j], input[i]
		}
	}
	input[0], input[j] = input[j], input[0]
	counts := j + 1

	if counts >= k {
		quickSelect(input[:j+1], k)
	} else {
		quickSelect(input[j+1:], k-counts)
	}
}
