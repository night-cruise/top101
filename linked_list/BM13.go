package linkedlist

// 快慢指针 + 链表反转
func IsPail(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	head2 := ReverseList2(slow.Next)
	slow.Next = nil

	p, q := head, head2
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p, q = p.Next, q.Next
	}

	return true
}
