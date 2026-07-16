package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    midIndex := len(nums)/2

    var node *TreeNode = &TreeNode{
        Val: nums[midIndex],
        Left: sortedArrayToBST(nums[:midIndex]),
        Right: sortedArrayToBST(nums[midIndex+1:]),
    }

    return node
}