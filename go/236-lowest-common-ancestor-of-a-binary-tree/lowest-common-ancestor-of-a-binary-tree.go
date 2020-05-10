package main

import "fmt"

func main() {
	p := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}
	q := &TreeNode{
		Val:   5,
		Left:  &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: p,
		},
	}
	a := &TreeNode{
		Val:   3,
		Left: q ,
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{
				Val:   0,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   8,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Print(*lowestCommonAncestor(a , p , q))

}
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	if root == nil {
		return nil
	}
	left := lowestCommonAncestor(root.Left, p , q)
	right := lowestCommonAncestor(root.Right, p , q)
	if right !=nil && left != nil {
		return root
	}
	if right != nil {
		return right
	}
	return left
}
