package linkedlist

// 双指针
func FindFirstCommonNode1(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	if pHead1 == nil || pHead2 == nil {
		return nil
	}
	p1, p2 := pHead1, pHead2

	for p1 != p2 {
		if p1 == nil {
			p1 = pHead2
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = pHead1
		} else {
			p2 = p2.Next
		}
	}

	return p1
}

// 哈希表
func FindFirstCommonNode2(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	m := map[*ListNode]bool{}

	for pHead1 != nil {
		m[pHead1] = true
		pHead1 = pHead1.Next
	}
	for pHead2 != nil && !m[pHead2] {
		pHead2 = pHead2.Next
	}

	return pHead2
}
