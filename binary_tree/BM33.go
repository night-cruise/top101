package binarytree

// Mirror(root) = root.Left->Mirror(root.Right) + root.Right->Mirror(root.Left)
func Mirror(pRoot *TreeNode) *TreeNode {
	if pRoot == nil {
		return nil
	}
	pRoot.Left, pRoot.Right = Mirror(pRoot.Right), Mirror(pRoot.Left)

	return pRoot
}
