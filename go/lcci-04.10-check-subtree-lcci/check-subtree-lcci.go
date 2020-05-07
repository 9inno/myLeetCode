package leetcode



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {

	if t1 == t2 {
		return true
	}

	if (t1==nil && t2!=nil) || (t1!=nil && t2==nil) {
		return false
	}
	if t1.Val != t2.Val {
		return checkSubTree(t1.Left,t2) || checkSubTree(t1.Right,t2)
	}
	return checkSubTree(t1.Left, t2.Left) || checkSubTree(t1.Right, t2.Right)

}