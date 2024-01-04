package heapstackqueue

// 小顶堆
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MinHeap) Less(i, j int) bool {
	return h[i] <= h[j]
}
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
