// link: https://leetcode.com/problems/binary-tree-postorder-traversal/

package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	var postorder func(*TreeNode)
	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}

		postorder(node.Left)
		postorder(node.Right)

		result = append(result, node.Val)

	}
	postorder(root)
	return result
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Right.Left = &TreeNode{Val: 3}

	fmt.Println(postorderTraversal(root))
}
