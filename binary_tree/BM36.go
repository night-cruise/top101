package binarytree

// 递归
//
// IsBalanced(root) = IsBalanced(root.Left) && IsBalanced(root.Right) && balanced(root)
func IsBalanced_Solution(pRoot *TreeNode) bool {
	var isBalancedHelper func(root *TreeNode) (int, bool)

	isBalancedHelper = func(root *TreeNode) (int, bool) {
		if root == nil {
			return 0, true
		}
		leftHeight, leftIsBalanced := isBalancedHelper(root.Left)
		rightHeight, rightIsBalanced := isBalancedHelper(root.Right)

		return Max(leftHeight, rightHeight) + 1, leftIsBalanced && rightIsBalanced && Absolute(leftHeight-rightHeight) <= 1
	}
	_, isBalanced := isBalancedHelper(pRoot)

	return isBalanced
}
