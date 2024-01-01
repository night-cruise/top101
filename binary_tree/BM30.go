package binarytree

// 中序遍历，设置一个前驱指针，在遍历的过程中设置前驱结点的 Right 指针和当前节点的 left 指针即可。
func Convert(pRootOfTree *TreeNode) *TreeNode {
	var prev, head *TreeNode
	var convertHelper func(*TreeNode)

	convertHelper = func(root *TreeNode) {
		if root == nil {
			return
		}
		convertHelper(root.Left)
		if head == nil {
			head = root
		}
		if prev != nil {
			prev.Right, root.Left = root, prev
		}
		prev = root
		convertHelper(root.Right)
	}
	convertHelper(pRootOfTree)

	return head
}
