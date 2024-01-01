package binarytree

// 层次遍历，在使用双队列的基础上再加一个变量表示当前访问的是奇数层还是偶数层，
// 如果是奇数层则从左到右访问节点的值，否则从右到左访问节点的值。
func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return [][]int{}
	}
	seq := [][]int{}
	q1, q2 := []*TreeNode{}, []*TreeNode{}

	q1 = append(q1, pRoot)
	i := 1
	for len(q1) > 0 || len(q2) > 0 {
		cur := []int{}
		if len(q1) > 0 {
			if i&1 == 1 {
				for i := 0; i < len(q1); i++ {
					cur = append(cur, q1[i].Val)
				}
			} else {
				for i := len(q1) - 1; i >= 0; i-- {
					cur = append(cur, q1[i].Val)
				}
			}
			for len(q1) > 0 {
				node := q1[0]
				if node.Left != nil {
					q2 = append(q2, node.Left)
				}
				if node.Right != nil {
					q2 = append(q2, node.Right)
				}
				q1 = q1[1:]
			}
		} else {
			if i&1 == 1 {
				for i := 0; i < len(q2); i++ {
					cur = append(cur, q2[i].Val)
				}
			} else {
				for i := len(q2) - 1; i >= 0; i-- {
					cur = append(cur, q2[i].Val)
				}
			}
			for len(q2) > 0 {
				node := q2[0]
				if node.Left != nil {
					q1 = append(q1, node.Left)
				}
				if node.Right != nil {
					q1 = append(q1, node.Right)
				}
				q2 = q2[1:]
			}
		}
		seq = append(seq, cur)
		i++
	}

	return seq
}
