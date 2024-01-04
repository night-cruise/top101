package heapstackqueue

// // 最小值栈
// var stack []int
// var minStack []int

// func PushBM43(node int) {
// 	stack = append(stack, node)
// 	if len(minStack) == 0 || node <= minStack[len(minStack)-1] {
// 		minStack = append(minStack, node)
// 	}
// }
// func PopBM43() {
// 	if stack[len(stack)-1] == minStack[len(minStack)-1] {
// 		minStack = minStack[:len(minStack)-1]
// 	}
// 	stack = stack[:len(stack)-1]
// }
// func TopBM43() int {
// 	return stack[len(stack)-1]
// }
// func MinBM43() int {
// 	return minStack[len(minStack)-1]
// }

// 元组，元组中的第二项保存当前最小元素值
var stack = [][]int{}

func PushBM43(node int) {
	if len(stack) == 0 || node < stack[len(stack)-1][1] {
		stack = append(stack, []int{node, node})
	} else {
		stack = append(stack, []int{node, stack[len(stack)-1][1]})
	}
}
func PopBM43() {
	stack = stack[:len(stack)-1]
}
func TopBM43() int {
	return stack[len(stack)-1][0]
}
func MinBM43() int {
	return stack[len(stack)-1][1]
}
