package binarytree

// BFS
//
// 使用一个前驱指针指向上一个遍历到的节点，如果当前节点不为空而前驱指针为空则说明不是完全二叉树
func IsCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	prev := &TreeNode{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if prev == nil && node != nil {
			return false
		}
		if node != nil {
			q = append(q, node.Left)
			q = append(q, node.Right)
		}
		prev = node
	}

	return true
}
