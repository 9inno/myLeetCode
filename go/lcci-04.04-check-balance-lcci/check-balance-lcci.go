package leetcode

import (
	"fmt"
	"math"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	result := true
	recursion(root, 0, &result)
	return result
}

func recursion(root *TreeNode, number int, result *bool) int {

	if root == nil || *result == false{
		return 0
	}

	left := recursion(root.Left, number + 1, result)
	right := recursion(root.Right, number + 1, result)
	fmt.Printf("left %d , right %d , val %d", left, right, root.Val)
	if math.Abs(float64(left - right)) > 1 {
		*result = false
	}
	return max(left, right) + 1
}

func max(a , b int) int {
	if a < b {
		return b
	}
	return a
}