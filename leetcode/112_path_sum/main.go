package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	remainder := targetSum - root.Val

	if remainder == 0 && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Left, remainder) || hasPathSum(root.Right, remainder)
}
