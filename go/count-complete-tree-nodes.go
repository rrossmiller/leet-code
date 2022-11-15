/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// make tree
	four := TreeNode{Val: 4}
	left := TreeNode{Val: 2, Left: &four}
	right := TreeNode{Val: 3}
	root := TreeNode{1, &left, &right}
	ans := countNodes(&root)
	fmt.Println(ans)
}

func bfs(root *TreeNode) (n int) {
	n = 0
	q := []*TreeNode{root}
	var curr *TreeNode

	for len(q) > 0 {
		n++
		curr = q[len(q)-1]
		q = q[:len(q)-1]

		if curr.Left != nil {
			q = append(q, curr.Left)
		}

		if curr.Right != nil {
			q = append(q, curr.Right)
		}
	}
	return
}

// the tree is full except for maybe the last layer
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return bfs(root)
}
