package heapstackqueue

// 暴力解，每次 Insert 时都对数据流进行排序
import (
	"container/heap"
	"sort"
)

var stream = []int{}

func Insert1(num int) {
	stream = append(stream, num)
	sort.Slice(stream, func(i, j int) bool { return stream[i] < stream[j] })
}

func GetMedian1() float64 {
	n := len(stream)
	if len(stream)&1 > 0 {
		return float64(stream[n/2])
	}
	return float64(stream[n/2-1]+stream[n/2]) / 2.0
}

// 对顶堆，使用一个大顶堆存储较小的一半元素，使用一个小顶堆存储较大的一半元素，
// 当数量为奇数时，大顶堆的堆顶元素即为中位数，否则大顶堆和小顶堆的堆顶元素的均值即为中位数。
var maxHeap = &MaxHeap{}
var minHeap = &MinHeap{}

func Insert2(num int) {
	if maxHeap.Len() == 0 || num <= (*maxHeap)[0] {
		heap.Push(maxHeap, num)
	} else {
		heap.Push(minHeap, num)
	}
	for maxHeap.Len()-minHeap.Len() > 1 {
		heap.Push(minHeap, heap.Pop(maxHeap))
	}
	for minHeap.Len() > maxHeap.Len() {
		heap.Push(maxHeap, heap.Pop(minHeap))
	}
}

func GetMedian2() float64 {
	n := maxHeap.Len() + minHeap.Len()
	if n&1 > 0 {
		return float64((*maxHeap)[0])
	}
	return float64((*maxHeap)[0]+(*minHeap)[0]) / 2.0
}
