package binarytree

// 层次遍历
//
// 用一个变量记录当前在第几层，如果是奇数层则从左到右遍历，否则从右到左遍历
func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return [][]int{}
	}
	res := [][]int{}
	q := []*TreeNode{pRoot}
	level := 1

	for len(q) > 0 {
		size := len(q)
		tmp := []int{}
		for i := 0; i < size; i++ {
			if level&1 > 0 {
				tmp = append(tmp, q[i].Val)
			} else {
				tmp = append(tmp, q[size-i-1].Val)
			}
		}
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, tmp)
		level++
	}

	return res
}
