package leetcode

import (
	"math"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return recursion(root, math.MaxInt64, math.MinInt64)
}

func recursion(root *TreeNode, max, min int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return recursion(root.Left, root.Val, min) && recursion(root.Right, max, root.Val)
}