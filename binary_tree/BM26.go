package binarytree

// 层次遍历
//
// 使用两个队列，队列 q1 表示正在访问的节点，队列 a2 表示下一层的节点，
// 当 q1 访问完成之后，q2 保存了下层的所有节点，当下一次循环的时候 q2 就表示正在访问的节点，q1 表示下一层的节点。
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	seq := [][]int{}
	q1, q2 := []*TreeNode{}, []*TreeNode{}

	q1 = append(q1, root)
	for len(q1) > 0 || len(q2) > 0 {
		cur := []int{}
		if len(q1) > 0 {
			for len(q1) > 0 {
				node := q1[0]
				cur = append(cur, node.Val)
				if node.Left != nil {
					q2 = append(q2, node.Left)
				}
				if node.Right != nil {
					q2 = append(q2, node.Right)
				}
				q1 = q1[1:]
			}
		} else {
			for len(q2) > 0 {
				node := q2[0]
				cur = append(cur, node.Val)
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
	}

	return seq
}
