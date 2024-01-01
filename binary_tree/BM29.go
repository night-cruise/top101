package binarytree

// 递归
//
// HasPathSum(root) = HasPathSum(root.Left, sum - root.Val) || HasPathSum(root.Right, sum - root.Val)
func HasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}
	return HasPathSum(root.Left, sum-root.Val) || HasPathSum(root.Right, sum-root.Val)
}
