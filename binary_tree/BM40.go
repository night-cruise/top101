package binarytree

// 递归
func ReConstructBinaryTree(preOrder []int, vinOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	rootIndex := 0
	for vinOrder[rootIndex] != preOrder[0] {
		rootIndex++
	}
	leftNumber := rootIndex - 0
	root := &TreeNode{Val: preOrder[0]}

	root.Left = ReConstructBinaryTree(preOrder[1:leftNumber+1], vinOrder[:rootIndex])
	root.Right = ReConstructBinaryTree(preOrder[leftNumber+1:], vinOrder[rootIndex+1:])

	return root
}
