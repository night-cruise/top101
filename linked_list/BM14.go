package linkedlist

// 双指针
func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy1, dummy2 := &ListNode{}, &ListNode{}
	tail1, tail2 := dummy1, dummy2

	i := 1
	for head != nil {
		if i&1 > 0 {
			tail1.Next = head
			tail1 = head
		} else {
			tail2.Next = head
			tail2 = head
		}
		i++
		head = head.Next
	}
	tail2.Next = nil
	tail1.Next = dummy2.Next

	return dummy1.Next
}
