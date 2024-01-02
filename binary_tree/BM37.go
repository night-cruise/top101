package binarytree

// 递归
//
// LCA(root, p, q) = judge(root.Val, p.Val, q.Val) + LCA(root.Left, p, q) + LCA(root.Right, p, q)
func LowestCommonAncestorBM37(root *TreeNode, p int, q int) int {
	if p <= root.Val && root.Val <= q || q <= root.Val && root.Val <= p {
		return root.Val
	}
	if p > root.Val {
		return LowestCommonAncestorBM37(root.Right, p, q)
	}
	return LowestCommonAncestorBM37(root.Left, p, q)
}
