package linkedlist

// 双指针 + 辅助数组
func AddInLis1(head1 *ListNode, head2 *ListNode) *ListNode {
	v1, v2 := []*ListNode{}, []*ListNode{}

	for p := head1; p != nil; p = p.Next {
		v1 = append(v1, p)
	}
	for p := head2; p != nil; p = p.Next {
		v2 = append(v2, p)
	}

	var newHead *ListNode
	p1, p2, tmp := len(v1)-1, len(v2)-1, 0

	for p1 >= 0 || p2 >= 0 {
		adds := tmp
		if p1 >= 0 {
			adds += v1[p1].Val
			p1--
		}
		if p2 >= 0 {
			adds += v2[p2].Val
			p2--
		}
		if adds >= 10 {
			adds -= 10
			tmp = 1
		} else {
			tmp = 0
		}
		node := &ListNode{Val: adds, Next: newHead}
		newHead = node
	}
	if tmp != 0 {
		node := &ListNode{Val: 1, Next: newHead}
		newHead = node
	}

	return newHead
}

// 双指针 + 反转链表
func AddInLis2(head1 *ListNode, head2 *ListNode) *ListNode {
	head1 = ReverseList2(head1)
	head2 = ReverseList2(head2)

	var newHead *ListNode
	var tmp int

	for head1 != nil || head2 != nil {
		node := &ListNode{Next: newHead}
		newHead = node

		add := tmp
		if head1 != nil {
			add += head1.Val
			head1 = head1.Next
		}
		if head2 != nil {
			add += head2.Val
			head2 = head2.Next
		}
		if add >= 10 {
			tmp = 1
			newHead.Val = add - 10
		} else {
			tmp = 0
			newHead.Val = add
		}
	}
	if tmp != 0 {
		node := &ListNode{Val: 1, Next: newHead}
		newHead = node
	}

	return newHead
}
