package binarytree

import (
	"log"
	"testing"
)

func TestPreOrder(t *testing.T) {
	// result F B A O D C E G I H
	result := preOrderBinaryTreeTraverse(TreeString)
	log.Println(result)
}

func preOrderBinaryTreeTraverse(node *TreeNode) (result []string) {
	stack := make([]*TreeNode, 0, 100)

	for node != nil || len(stack) > 0 {
		if node != nil {
			result = append(result, node.Val)
			stack = append(stack, node)
			node = node.Left
		} else {
			topOfStack := len(stack) - 1
			node = stack[topOfStack]
			stack = stack[:topOfStack]
			node = node.Right
		}
	}

	return
}
