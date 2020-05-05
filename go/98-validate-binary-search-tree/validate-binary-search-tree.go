package main

import "math"

func main() {
	a := TreeNode{
		Val:   1,
		Left:  &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	isValidBST(&a)
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return recursion(root, math.MinInt64, math.MaxInt64)
}

func recursion(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return recursion(root.Left, lower, root.Val) && recursion(root.Right, root.Val, upper)
}