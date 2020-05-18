package leetcode




type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		var tmpQueue []*TreeNode
		var tmpResult []int
		for _ , value := range queue {
			tmpResult = append(tmpResult, value.Val)
			if value.Left != nil {
				tmpQueue = append(tmpQueue, value.Left)
			}
			if value.Right != nil {
				tmpQueue = append(tmpQueue, value.Right)
			}
		}
		result = append(result, tmpResult)
		queue = tmpQueue
	}
	return result
}