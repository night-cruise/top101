package binarytree

// 递归（O(N^2)）
//
// LCA(root, o1, o2) = judge(root.Val, o1, o2) + LCA(root.Left, o1, o2) + LCA(root.Right, o1, o2)
func LowestCommonAncestorBM38_1(root *TreeNode, o1 int, o2 int) int {
	if root.Val == o1 || root.Val == o2 {
		return root.Val
	}
	o1InLeft, o2InLeft := LCAHelper(root.Left, o1), LCAHelper(root.Left, o2)

	if o1InLeft && !o2InLeft || !o1InLeft && o2InLeft {
		return root.Val
	}
	if o1InLeft {
		return LowestCommonAncestorBM38_1(root.Left, o1, o2)
	}
	return LowestCommonAncestorBM38_1(root.Right, o1, o2)
}

func LCAHelper(root *TreeNode, o int) bool {
	if root == nil {
		return false
	}
	return root.Val == o || LCAHelper(root.Left, o) || LCAHelper(root.Right, o)
}

// 回溯
//
// 使用一个额外的数组存储搜索路径
func LowestCommonAncestorBM38_2(root *TreeNode, o1 int, o2 int) int {
	if root.Val == o1 || root.Val == o2 {
		return root.Val
	}
	o1Path, o2Path := SearchPath(root, o1), SearchPath(root, o2)

	var i int
	for i = 1; i < len(o1Path) && i < len(o2Path); i++ {
		if o1Path[i] != o2Path[i] {
			break
		}
	}

	return o1Path[i-1]
}

func SearchPath(root *TreeNode, o int) []int {
	path := []int{}

	var searchPathHelper func(root *TreeNode, o int) bool

	searchPathHelper = func(root *TreeNode, o int) bool {
		if root == nil {
			return false
		}
		path = append(path, root.Val)
		if root.Val == o || searchPathHelper(root.Left, o) || searchPathHelper(root.Right, o) {
			return true
		}
		path = path[:len(path)-1]
		return false
	}
	searchPathHelper(root, o)

	return path
}

// 递归
//
// LCA(root, o1, o2) = judge(root.Val, o1, o2) + LCA(root.Left, o1, o2) + LCA(root.Right, o1, o2)
func LowestCommonAncestorBM38_3(root *TreeNode, o1 int, o2 int) int {
	if root == nil {
		return -1
	}
	if root.Val == o1 || root.Val == o2 {
		return root.Val
	}
	lcaLeft, lcaRight := LowestCommonAncestorBM38_3(root.Left, o1, o2), LowestCommonAncestorBM38_3(root.Right, o1, o2)
	if lcaLeft == -1 {
		return lcaRight
	}
	if lcaRight == -1 {
		return lcaLeft
	}
	return root.Val
}
