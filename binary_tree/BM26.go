package binarytree

// 层序遍历
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	seq := [][]int{}
	q := []*TreeNode{}

	q = append(q, root)
	for len(q) > 0 {
		n := len(q)
		cur := []int{}
		for i := 0; i < n; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
			cur = append(cur, node.Val)
		}
		seq = append(seq, cur)
	}

	return seq
}
