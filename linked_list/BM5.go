package linkedlist

import "container/heap"

// 两两合并
func MergeKLists1(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	head := lists[0]
	for i := 1; i < len(lists); i++ {
		head = mergeTowList(head, lists[i])
	}
	return head
}

func mergeTowList(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}

	if pHead1.Val <= pHead2.Val {
		pHead1.Next = mergeTowList(pHead1.Next, pHead2)
		return pHead1
	}
	pHead2.Next = mergeTowList(pHead1, pHead2.Next)
	return pHead2
}

// 分治
func MergeKLists2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	leftHead := MergeKLists2(lists[0 : len(lists)/2])
	rightHead := MergeKLists2(lists[len(lists)/2:])

	return mergeTowList(leftHead, rightHead)
}

// 优先队列
type BinaryHeap []*ListNode

func (h BinaryHeap) Len() int {
	return len(h)
}
func (h BinaryHeap) Less(i, j int) bool {
	return h[i].Val > h[j].Val
}
func (h BinaryHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *BinaryHeap) Push(x any) {
	*h = append(*h, x.(*ListNode))
}
func (h *BinaryHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

func MergeKLists3(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := &BinaryHeap{}

	for _, list := range lists {
		p := list
		for p != nil {
			heap.Push(h, p)
			p = p.Next
		}
	}
	var newHead *ListNode
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		node.Next = newHead
		newHead = node
	}

	return newHead
}
