package linkedlist

// 快慢指针
func EntryNodeOfLoop(pHead *ListNode) *ListNode {
	slow, fast := pHead, pHead

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}

	slow = pHead
	for slow != fast {
		slow, fast = slow.Next, fast.Next
	}

	return slow
}
