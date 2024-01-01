package linkedlist

// 快慢指针
func HasCycle1(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow, fast := slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}

	return fast != nil && fast.Next != nil
}

// 哈希表
func HasCycle2(head *ListNode) bool {
	m := map[*ListNode]bool{}

	for head != nil && !m[head] {
		m[head] = true
		head = head.Next
	}

	return head != nil
}
