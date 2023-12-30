package linkedlist

// 双指针
func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	var h1, h2, p1, p2 *ListNode

	i := 1
	for p := head; p != nil; p = p.Next {
		if i&1 > 0 {
			if h1 == nil {
				h1, p1 = p, p
			} else {
				p1.Next = p
				p1 = p
			}
		} else {
			if h2 == nil {
				h2, p2 = p, p
			} else {
				p2.Next = p
				p2 = p
			}
		}
		i++
	}
	p1.Next, p2.Next = h2, nil

	return h1
}
