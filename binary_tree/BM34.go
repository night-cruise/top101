package binarytree

// 中序遍历，设置前驱节点
func IsValidBST1(root *TreeNode) bool {
	var prev *TreeNode
	isValidBST := true
	var inorderTraversal func(root *TreeNode)

	inorderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraversal(root.Left)
		if prev != nil && prev.Val > root.Val {
			isValidBST = false
		}
		prev = root
		inorderTraversal(root.Right)
	}
	inorderTraversal(root)

	return isValidBST
}

// 递归
//
// IsValidBST(root) = IsValidBST(root.Left) && IsValidBST(root.Right) && root.Left.val < root.val < root.Right.Val
func IsValidBST2(root *TreeNode) bool {
	var isValidBSTHelper func(root, parent *TreeNode) bool

	isValidBSTHelper = func(root, parent *TreeNode) bool {
		if root == nil {
			return true
		}
		if root.Left != nil {
			if root.Left.Val > root.Val {
				return false
			}
			if parent != nil && parent.Right == root && root.Left.Val < parent.Val {
				return false
			}
		}
		if root.Right != nil {
			if root.Right.Val < root.Val {
				return false
			}
			if parent != nil && parent.Left == root && root.Right.Val > parent.Val {
				return false
			}
		}
		return isValidBSTHelper(root.Left, root) && isValidBSTHelper(root.Right, root)
	}

	return isValidBSTHelper(root, nil)
}
