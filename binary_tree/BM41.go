package binarytree

func Solve(preOrder []int, inOrder []int) []int {
	_ = rightView1(buildTree(preOrder, inOrder))
	return rightView2(buildTree(preOrder, inOrder))
}

// DFS
func buildTree(preOrder, inOrder []int) *TreeNode {
	if len(preOrder) == 0 {
		return nil
	}
	var rootIndex int
	for rootIndex = 0; rootIndex < len(inOrder); rootIndex++ {
		if inOrder[rootIndex] == preOrder[0] {
			break
		}
	}
	root := &TreeNode{Val: preOrder[0]}
	root.Left = buildTree(preOrder[1:1+rootIndex], inOrder[:rootIndex])
	root.Right = buildTree(preOrder[1+rootIndex:], inOrder[rootIndex+1:])

	return root
}

// BFS
func rightView1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	q := []*TreeNode{root}

	for len(q) > 0 {
		size := len(q)
		res = append(res, q[size-1].Val)
		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}

	return res
}

// DFS
func rightView2(root *TreeNode) []int {
	m := map[int]int{}
	var rightViewHelper func(root *TreeNode, level int, m map[int]int)

	rightViewHelper = func(root *TreeNode, level int, m map[int]int) {
		if root == nil {
			return
		}
		m[level] = root.Val
		rightViewHelper(root.Left, level+1, m)
		rightViewHelper(root.Right, level+1, m)
	}
	rightViewHelper(root, 0, m)

	res := []int{}

	for level := 0; ; level++ {
		val, ok := m[level]
		if !ok {
			break
		}
		res = append(res, val)
	}

	return res
}
