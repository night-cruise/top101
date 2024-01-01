package binarytree

// 递归
//
// depth(root) = max(depth(root.Left), depth(root.Right)) + 1
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return Max(MaxDepth(root.Left), MaxDepth(root.Right)) + 1
}
