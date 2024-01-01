package linkedlist

// 两次遍历
func FindKthToTail1(pHead *ListNode, k int) *ListNode {
	n := getLength(pHead)

	if k > n {
		return nil
	}
	for i := 0; i < n-k; i++ {
		pHead = pHead.Next
	}

	return pHead
}

// 快慢指针
//
// 快指针先走 k 步，然后让快慢指针同时走，当快指针走到末尾的时候，慢指针指向倒数第 k 个节点
func FindKthToTail2(pHead *ListNode, k int) *ListNode {
	slow, fast := pHead, pHead

	var i int
	for i = 0; i < k && fast != nil; i++ {
		fast = fast.Next
	}
	// 链表的长度小于 k
	if i < k {
		return nil
	}
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}

	return slow
}
