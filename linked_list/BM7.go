package linkedlist

// 快慢指针
func EntryNodeOfLoop1(pHead *ListNode) *ListNode {
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

// 哈希表
func EntryNodeOfLoop2(pHead *ListNode) *ListNode {
	m := map[*ListNode]bool{}

	for pHead != nil && !m[pHead] {
		m[pHead] = true
		pHead = pHead.Next
	}

	return pHead
}
