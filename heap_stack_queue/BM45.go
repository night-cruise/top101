package heapstackqueue

import "container/heap"

// 暴力枚举
func MaxInWindows1(num []int, size int) []int {
	if size > len(num) || size == 0 {
		return []int{}
	}
	left, right := 0, size-1
	res := []int{}

	for right < len(num) {
		maxVal := num[left]
		for i := left + 1; i <= right; i++ {
			if num[i] > maxVal {
				maxVal = num[i]
			}
		}
		res = append(res, maxVal)
		left++
		right++
	}

	return res
}

// 大顶堆
type PriorityQueue [][]int

func (q PriorityQueue) Len() int {
	return len(q)
}
func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}
func (q PriorityQueue) Less(i, j int) bool {
	return q[i][0] >= q[j][0]
}
func (q *PriorityQueue) Push(x any) {
	*q = append(*q, x.([]int))
}
func (q *PriorityQueue) Pop() any {
	x := (*q)[len(*q)-1]
	*q = (*q)[:len(*q)-1]

	return x
}

func MaxInWindows2(num []int, size int) []int {
	if size > len(num) || size == 0 {
		return []int{}
	}
	q := &PriorityQueue{}
	for i := 0; i < size; i++ {
		heap.Push(q, []int{num[i], i})
	}
	left, right := 0, size-1
	res := []int{}

	for right < len(num) {
		res = append(res, (*q)[0][0])
		left++
		right++
		if right >= len(num) {
			break
		}
		for q.Len() > 0 && (*q)[0][1] < left {
			heap.Pop(q)
		}
		heap.Push(q, []int{num[right], right})
	}

	return res
}

// 单调递减队列
func MaxInWindows3(num []int, size int) []int {
	if size > len(num) || size == 0 {
		return []int{}
	}
	q, res := []int{}, []int{}
	for i := 0; i < size; i++ {
		for len(q) > 0 && num[q[len(q)-1]] < num[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	left, right := 0, size-1

	for right < len(num) {
		res = append(res, num[q[0]])
		left++
		right++
		if right >= len(num) {
			continue
		}
		for len(q) > 0 && q[0] < left {
			q = q[1:]
		}
		for len(q) > 0 && num[q[len(q)-1]] < num[right] {
			q = q[:len(q)-1]
		}
		q = append(q, right)
	}

	return res
}
