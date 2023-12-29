package linkedlist

// 双指针
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	q, p := dummy, dummy.Next
	length := getLength(head)

	for i := 0; i < length-n; i++ {
		q, p = p, p.Next
	}
	q.Next = p.Next

	return dummy.Next
}
