package binarytree

type TreeNodeInt struct {
	Val   int
	Left  *TreeNodeInt
	Right *TreeNodeInt
}

type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

var TreeString = &TreeNode{
	Val: "F",
	Right: &TreeNode{
		Val: "G",
		Right: &TreeNode{
			Val: "I",
			Left: &TreeNode{
				Val: "H",
			},
		},
	},
	Left: &TreeNode{
		Val: "B",
		Left: &TreeNode{
			Val: "A",
			Right: &TreeNode{
				Val: "O",
			},
		},
		Right: &TreeNode{
			Val: "D",
			Left: &TreeNode{
				Val: "C",
			},
			Right: &TreeNode{
				Val: "E",
			},
		},
	},
}
