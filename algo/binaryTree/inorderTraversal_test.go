package binarytree

import (
	"fmt"
	"testing"
)

func TestInorder(t *testing.T) {
	root := &TreeNodeInt{
		Val: 1,
		Right: &TreeNodeInt{
			Val:  2,
			Left: &TreeNodeInt{Val: 3},
		},
	}

	res := inorderTraversal(root)
	fmt.Println(res)
	t.Error(1)
}

func inorderTraversal(root *TreeNodeInt) []int {
	// stack ← empty stack

	// while not stack.isEmpty() or node ≠ null
	//     if node ≠ null
	//         stack.push(node)
	//         node ← node.left
	//     else
	//         node ← stack.pop()
	//         visit(node)
	//         node ← node.right

	res := make([]int, 0)
	stack := make([]*TreeNodeInt, 0, 100)
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			res = append(res, root.Val)
			root = root.Right
		}
	}

	return res

}
