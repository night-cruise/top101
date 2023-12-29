package linkedlist

// 双链表
func ReverseList1(head *ListNode) *ListNode {
	var newHead *ListNode

	for head != nil {
		p := head.Next
		head.Next = newHead
		newHead = head
		head = p
	}

	return newHead
}

// 递归
func ReverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := ReverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
