package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root != nil {
		l, r := 1, 1
		l += maxDepth(root.Left)
		r += maxDepth(root.Right)

		if l > r {
			return l
		}
		return r
	}

	return 0
}
