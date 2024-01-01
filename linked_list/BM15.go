package linkedlist

// 双指针
func DeleteDuplicatesBM15(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev, cur := head, head.Next

	for cur != nil {
		if prev.Val == cur.Val {
			prev.Next = cur.Next
			cur = prev.Next
		} else {
			prev, cur = cur, cur.Next
		}
	}

	return head
}
