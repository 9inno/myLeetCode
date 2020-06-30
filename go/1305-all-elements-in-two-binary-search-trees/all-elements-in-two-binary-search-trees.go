package leetcode

import "sort"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	result := []int{}
	dfs(root1 , & result)
	dfs(root2 , & result)
	sort.Ints(result)
	return result
}

func dfs(root *TreeNode , result *[]int)  {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	dfs(root.Left, result)
	dfs(root.Right, result)
}