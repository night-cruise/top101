package binarytree

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	seq := []int{root.Val}
	seq = append(seq, PreorderTraversal(root.Left)...)
	seq = append(seq, PreorderTraversal(root.Right)...)

	return seq
}
