package linkedlist

// 双指针
func DeleteDuplicatesBM15(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 101, Next: head}
	q, p := dummy, dummy.Next

	for p != nil {
		if q.Val == p.Val {
			q.Next = p.Next
			p = q.Next
		} else {
			q, p = p, p.Next
		}
	}

	return dummy.Next
}
