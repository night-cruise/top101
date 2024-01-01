package binarytree

func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	seq := InorderTraversal(root.Left)
	seq = append(seq, root.Val)
	seq = append(seq, InorderTraversal(root.Right)...)

	return seq
}
