package binarytree

func PostorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	seq := PostorderTraversal(root.Left)
	seq = append(seq, PostorderTraversal(root.Right)...)
	seq = append(seq, root.Val)

	return seq
}
