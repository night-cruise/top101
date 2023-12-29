package linkedlist

// 快慢指针
func HasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast := slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}

	return fast != nil && fast.Next != nil
}
