package heapstackqueue

import "container/heap"

func FindKth1(a []int, n int, K int) int {
	h := &MinHeap{}

	for _, item := range a {
		heap.Push(h, item)
		if h.Len() > K {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)
}

// 快速选择
func FindKth2(a []int, n int, K int) int {
	return quickSelectFindKth(a, K)
}

func quickSelectFindKth(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	i, j, privot := -1, len(nums), nums[0]

	for i < j {
		for {
			i++
			if nums[i] <= privot {
				break
			}
		}
		for {
			j--
			if nums[j] >= privot {
				break
			}
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[0], nums[j] = nums[j], nums[0]
	counts := j + 1

	if counts >= k {
		return quickSelectFindKth(nums[:j+1], k)
	} else {
		return quickSelectFindKth(nums[j+1:], k-counts)
	}
}
