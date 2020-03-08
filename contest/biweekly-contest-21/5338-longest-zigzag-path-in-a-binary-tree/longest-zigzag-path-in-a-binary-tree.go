package leetcode


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func longestZigZag(root *TreeNode) int {
	maxPath := 0
	recursion(root.Left, 1, "left", &maxPath)
	recursion(root.Right, 1, "right", &maxPath)
	return maxPath
}

func recursion(root *TreeNode, count int , direction string, maxPath *int)  {
	if *maxPath < count - 1 {
		*maxPath = count - 1
	}
	if root == nil {
		return
	}
	if direction == "right" {
		recursion(root.Left, count + 1, "left", maxPath)
		recursion(root.Right, 1 , "right", maxPath)
	} else  {
		recursion(root.Right, count + 1, "right",maxPath)
		recursion(root.Left, 1 , "left",maxPath)
	}
}