package linkedlist

// 遍历，使用虚拟节点减少特殊情况的考虑
func ReverseKGroup1(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy, length := &ListNode{Next: head}, getLength(head)

	var preNode, h, p, q *ListNode
	preNode = dummy

	for i := 0; i < length/k; i++ {
		h = preNode.Next
		q, p = h, h.Next
		for j := 0; j < k-1; j++ {
			tmp := p.Next
			p.Next = q
			q, p = p, tmp
		}
		preNode.Next, h.Next = q, p
		preNode = h
	}

	return dummy.Next
}

func getLength(head *ListNode) int {
	length := 0

	for head != nil {
		length++
		head = head.Next
	}

	return length
}

// 递归
//
// Reverse(head, k) = reverse(head..head+k).Tail -> Reverse(head+k, k).Head
func ReverseKGroup2(head *ListNode, k int) *ListNode {
	tail := head
	for i := 0; i < k-1; i++ {
		if tail == nil {
			break
		}
		tail = tail.Next
	}
	if tail == nil {
		return head
	}
	head2 := tail.Next
	tail.Next = nil
	reverse(head)
	head.Next = ReverseKGroup2(head2, k)

	return tail
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead
}
