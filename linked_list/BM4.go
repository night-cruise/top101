package linkedlist

// 双指针
func Merge1(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	var p, q, k, newHead *ListNode

	p, q = pHead1, pHead2
	for p != nil && q != nil {
		var minNode *ListNode
		if p.Val <= q.Val {
			minNode, p = p, p.Next
		} else {
			minNode, q = q, q.Next
		}
		if k == nil {
			k, newHead = minNode, minNode
		} else {
			k.Next = minNode
			k = k.Next
		}
	}

	if p != nil {
		k.Next = p
	} else {
		k.Next = q
	}

	return newHead
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
