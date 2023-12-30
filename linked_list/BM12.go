package linkedlist

import "sort"

// 分治 + 有序链表合并 + 快慢指针
func SortInList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	head2 := slow.Next
	slow.Next = nil

	head1, head2 := SortInList1(head), SortInList1(head2)
	var newHead, tail *ListNode

	for head1 != nil && head2 != nil {
		var node *ListNode
		if head1.Val <= head2.Val {
			node = head1
			head1 = head1.Next
		} else {
			node = head2
			head2 = head2.Next
		}
		if tail == nil {
			newHead, tail = node, node
		} else {
			tail.Next = node
			tail = node
		}
	}
	if head1 != nil {
		tail.Next = head1
	} else {
		tail.Next = head2
	}

	return newHead
}

// type NodeList []*ListNode

// func (list NodeList) Len() int {
// 	return len(list)
// }
// func (list NodeList) Swap(i, j int) {
// 	list[i], list[j] = list[j], list[i]
// }
// func (list NodeList) Less(i, j int) bool {
// 	return list[i].Val > list[j].Val
// }

// 辅助数组
func SortInList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	v := []*ListNode{}

	for p := head; p != nil; p = p.Next {
		v = append(v, p)
	}

	sort.Slice(v, func(i, j int) bool { return v[i].Val > v[j].Val })

	var newHead *ListNode
	for _, node := range v {
		node.Next = newHead
		newHead = node
	}

	return newHead
}
