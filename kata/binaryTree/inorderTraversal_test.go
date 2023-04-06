package main

import (
	"log"
	"testing"
)

func TestInorderBinaryTreeTraverse(t *testing.T) {
	result := inorderTraversal(Tree)
	log.Println(result)
}

func inorderTraversal(root *TreeNode) (result []string) {
	stack := make([]*TreeNode, 0, 100)
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			end := len(stack) - 1
			root = stack[end]
			result = append(result, root.Val)
			stack = stack[:end]
			root = root.Right
		}
	}
	return
}
