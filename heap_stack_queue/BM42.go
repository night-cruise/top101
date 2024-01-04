package heapstackqueue

var stack1 []int
var stack2 []int

func PushBM42(node int) {
	stack1 = append(stack1, node)
}

func PopBM42() int {
	if len(stack2) == 0 {
		for len(stack1) > 0 {
			node := stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
			stack2 = append(stack2, node)
		}
	}
	node := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]

	return node
}
