package main

func main() {
	p := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	a := &TreeNode{
		Val:   2,
		Left: p,
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}


	inorderSuccessor(a ,p)
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root==nil||p==nil{
		return nil
	}
	if p.Val>=root.Val{
		return inorderSuccessor(root.Right,p)
	}else{
		left:=inorderSuccessor(root.Left,p)
		if left!=nil{
			return left
		}
	}
	return root

}