package main


func invertTree(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}

	return root
}

func invertTreeIter(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := New()
	stack.Push(root)

	for stack.Len() > 0 {
		node := stack.Pop()
		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}

	}

	return root
}

// https://pkg.go.dev/github.com/golang-collections/collections/stack
type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value *TreeNode
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}


// Pop the top item of the stack and return it
func (this *Stack) Pop() *TreeNode {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value *TreeNode) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}
