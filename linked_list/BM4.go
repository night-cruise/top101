package linkedlist

// 双指针，使用一个辅助节点能更简单
func Merge1(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	newHead := &ListNode{}
	tail := newHead

	for pHead1 != nil && pHead2 != nil {
		var minNode *ListNode
		if pHead1.Val < pHead2.Val {
			minNode, pHead1 = pHead1, pHead1.Next
		} else {
			minNode, pHead2 = pHead2, pHead2.Next
		}
		tail.Next = minNode
		tail = minNode
	}
	if pHead1 != nil {
		tail.Next = pHead1
	} else {
		tail.Next = pHead2
	}

	return newHead.Next
}

// 递归
func Merge2(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}

	if pHead1.Val <= pHead2.Val {
		pHead1.Next = Merge2(pHead1.Next, pHead2)
		return pHead1
	}
	pHead2.Next = Merge2(pHead1, pHead2.Next)
	return pHead2
}
