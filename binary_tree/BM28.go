package binarytree

// 递归
//
// depth(root) = max(depth(root.Left), depth(root.Right)) + 1
func MaxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(MaxDepth1(root.Left), MaxDepth1(root.Right)) + 1
}

// 层次遍历
func MaxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	level := 1

	for len(q) > 0 {
		size := len(q)
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
		level++
	}

	return level
}
