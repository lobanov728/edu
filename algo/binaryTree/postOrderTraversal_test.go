package binarytree

import (
	"log"
	"testing"
)

func TestPostOrder(t *testing.T) {
	// result O A C E D B H I G F
	result := postOrderBinaryTreeTraverse(TreeString)
	log.Println(result)
}

func postOrderBinaryTreeTraverse(node *TreeNode) (result []string) {
	stack := make([]*TreeNode, 0, 100)
	var lastNodeVisited *TreeNode
	for node != nil || len(stack) > 0 {
		// f
		// fb
		// fba
		// fbao
		// fba  o
		// fb   a
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			topOfStack := len(stack) - 1
			peekNode := stack[topOfStack]

			if peekNode.Right != nil && lastNodeVisited != peekNode.Right {
				node = peekNode.Right
			} else {
				result = append(result, peekNode.Val)
				lastNodeVisited = peekNode
				stack = stack[:topOfStack] // pop element
			}
		}
	}
	// procedure iterativePostorder(node)
	// stack ← empty stack
	// lastNodeVisited ← null
	// while not stack.isEmpty() or node ≠ null
	//     if node ≠ null
	//         stack.push(node)
	//         node ← node.left
	//     else
	//         peekNode ← stack.peek()
	//         // if right child exists and traversing node
	//         // from left child, then move right
	//         if peekNode.right ≠ null and lastNodeVisited ≠ peekNode.right
	//             node ← peekNode.right
	//         else
	//             visit(peekNode)
	//             lastNodeVisited ← stack.pop()

	return
}
