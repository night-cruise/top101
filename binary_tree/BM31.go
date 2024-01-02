package binarytree

// 递归
//
// IsSymmetrical(root1, root2) = root1.Val == root2.Val && IsSymmetrical(root1.Left, root2.Right) && IsSymmetrical(root1.Right, root2.Left)
func IsSymmetrical1(pRoot *TreeNode) bool {
	return pRoot == nil || IsSymmetricalHelper(pRoot.Left, pRoot.Right)
}

func IsSymmetricalHelper(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return IsSymmetricalHelper(root1.Left, root2.Right) && IsSymmetricalHelper(root1.Right, root2.Left)
}

// 迭代
func IsSymmetrical12(pRoot *TreeNode) bool {
	if pRoot == nil {
		return true
	}
	q := []*TreeNode{pRoot.Left, pRoot.Right}
	for len(q) > 0 {
		node1, node2 := q[0], q[1]
		q = q[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		q = append(q, node1.Left)
		q = append(q, node2.Right)
		q = append(q, node1.Right)
		q = append(q, node2.Left)
	}

	return true
}
