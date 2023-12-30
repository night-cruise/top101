package linkedlist

// 双指针
func DeleteDuplicatesBM16(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{Val: 1001, Next: head}
	q, p := dummy, dummy.Next

	for p != nil && p.Next != nil {
		if p.Val == p.Next.Val {
			for p != nil && p.Val == q.Next.Val {
				p = p.Next
			}
			q.Next = p
		} else {
			q, p = p, p.Next
		}
	}

	return dummy.Next
}
