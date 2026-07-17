package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var checkHeight func(node *TreeNode) int

	checkHeight = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftHeight := checkHeight(node.Left)

		if leftHeight == -1 {
			return -1
		}

		rightHeight := checkHeight(node.Right)

		if rightHeight == -1 {
			return -1
		}

		if abs(leftHeight-rightHeight) > 1 {
			return -1
		}

		return max(leftHeight, rightHeight) + 1
	}

	return checkHeight(root) != -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
