package binarytree

import (
	"fmt"
	"strconv"
	"strings"
)

// BFS
func Serialize1(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	res := strconv.Itoa(root.Val)
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node.Left == nil {
			res += ",#"
		} else {
			res += fmt.Sprintf(",%v", node.Left.Val)
			q = append(q, node.Left)
		}
		if node.Right == nil {
			res += ",#"
		} else {
			res += fmt.Sprintf(",%v", node.Right.Val)
			q = append(q, node.Right)
		}
	}

	return res
}

func Deserialize1(s string) *TreeNode {
	if s[0] == '#' {
		return nil
	}
	vals := strings.Split(s, ",")
	rootVal, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: rootVal}
	q := []*TreeNode{root}

	i := 1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if vals[i] != "#" {
			leftVal, _ := strconv.Atoi(vals[i])
			node.Left = &TreeNode{Val: leftVal}
			q = append(q, node.Left)
		}
		if vals[i+1] != "#" {
			rightVal, _ := strconv.Atoi(vals[i+1])
			node.Right = &TreeNode{Val: rightVal}
			q = append(q, node.Right)
		}

		i += 2
	}

	return root
}

// DFS
func Serializ2(root *TreeNode) string {
	var serializeHelper func(root *TreeNode) string

	serializeHelper = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}
		res := fmt.Sprintf("%v,", root.Val)
		res += serializeHelper(root.Left)
		res += ","
		res += serializeHelper(root.Right)

		return res
	}

	return serializeHelper(root)
}

func Deserialize2(s string) *TreeNode {
	// var deserializeHelper func(s []string) *TreeNode

	// deserializeHelper = func(s []string) *TreeNode {
	// 	if s[0] == "#" {
	// 		return nil
	// 	}
	// 	val, _ := strconv.Atoi(s[0])
	// 	root := &TreeNode{Val: val}

	// 	// 找到右子树的起始位置
	// 	i, need := 1, 1
	// 	for need > 0 {
	// 		if s[i] == "#" {
	// 			need--
	// 		} else {
	// 			need++
	// 		}
	// 		i++
	// 	}
	// 	root.Left = deserializeHelper(s[1:i])
	// 	root.Right = deserializeHelper(s[i:])

	// 	return root
	// }
	// return deserializeHelper(strings.Split(s, ","))

	// 更简单的写法
	var deserializeHelper func(s []string, i *int) *TreeNode

	deserializeHelper = func(s []string, i *int) *TreeNode {
		if s[*i] == "#" {
			*i += 1
			return nil
		}

		val, _ := strconv.Atoi(s[*i])
		root := &TreeNode{Val: val}

		*i += 1
		root.Left = deserializeHelper(s, i)
		root.Right = deserializeHelper(s, i)

		return root
	}

	i := 0

	return deserializeHelper(strings.Split(s, ","), &i)
}

// 1 2 # # 3 6 # # 7 # #
