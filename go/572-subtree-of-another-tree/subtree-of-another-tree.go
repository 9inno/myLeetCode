package leetcode



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s ==nil {
		return false
	}
	return recursion(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func recursion(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val == t.Val {
		return recursion(s.Left, t.Left) && recursion(s.Right, t.Right)
	}
	return false
}