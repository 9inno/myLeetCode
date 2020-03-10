package leetcode




type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxPath := 0
	recursion(root, &maxPath)
	return maxPath
}

func recursion(root *TreeNode, maxPath *int) int {
	if root == nil {
		return 0
	}
	left := recursion(root.Left, maxPath)
	right := recursion(root.Right, maxPath)
	if *maxPath < left + right {
		*maxPath = left + right
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

//需要注意 有 不经过根节点的可能
//[4,-7,-3,null,null,-9,-3,9,-7,-4,null,6,null,-6,-6,null,null,0,6,5,null,9,null,null,-1,-4,null,null,null,-2] 这样的树