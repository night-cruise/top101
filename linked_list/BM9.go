package linkedlist

// 双指针 + 遍历两次
func RemoveNthFromEnd1(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}

	q, p := dummy, dummy.Next
	length := getLength(head)

	for i := 0; i < length-n; i++ {
		q, p = p, p.Next
	}
	q.Next = p.Next

	return dummy.Next
}

// 快慢指针 + 虚拟节点
func RemoveNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	prev, slow, fast := dummy, head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		prev, slow, fast = slow, slow.Next, fast.Next
	}
	prev.Next = slow.Next

	return dummy.Next
}
