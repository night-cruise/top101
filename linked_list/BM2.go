package linkedlist

// 遍历，增加一个虚拟节点可以减少特殊情况的考虑
func ReverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	dummy := &ListNode{Next: head}

	preNode := dummy
	for i := 1; i < m; i++ {
		preNode = preNode.Next
	}
	h := preNode.Next
	q, p := h, h.Next

	for i := m; i < n; i++ {
		tmp := p.Next
		p.Next = q
		q, p = p, tmp
	}
	preNode.Next, h.Next = q, p

	return dummy.Next
}
