package linkedlist

// 遍历
func FindKthToTail(pHead *ListNode, k int) *ListNode {
	n := getLength(pHead)

	if k > n {
		return nil
	}
	for i := 0; i < n-k; i++ {
		pHead = pHead.Next
	}

	return pHead
}
