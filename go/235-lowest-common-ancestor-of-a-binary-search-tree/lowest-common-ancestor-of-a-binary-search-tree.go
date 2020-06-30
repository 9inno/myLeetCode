package main

func main() {
	p := &TreeNode{
		Val:   4,
		Left:  &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   5,
			Left:  nil,
			Right: nil,
		},
	}
	q := &TreeNode{
		Val:   2,
		Left:  &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: p,
	}

	a := &TreeNode{
		Val:   6,
		Left: q  ,
		Right: &TreeNode{
			Val:   8,
			Left:  &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}

	lowestCommonAncestor(a,p, q )
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p , q)
	right := lowestCommonAncestor(root.Left, p , q)
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}