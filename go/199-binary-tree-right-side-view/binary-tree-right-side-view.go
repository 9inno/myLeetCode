package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var tmpQueue []*TreeNode
		result = append(result, queue[len(queue) - 1].Val)
		for _ , node := range queue {
			if node.Left != nil {
				tmpQueue = append(tmpQueue , node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue , node.Right)
			}
		}
		queue = tmpQueue
	}
	return result
}



func rightSideView1(root *TreeNode) []int {
	var recursion func(node *TreeNode, deep int)
	var result []int
	recursion = func(node *TreeNode, deep int) {
		if node == nil {
			return
		}
		if deep == len(result){
			result = append(result, node.Val)
		}
		recursion(node.Right , deep + 1)
		recursion(node.Left , deep + 1)
	}
	recursion(root , 0)
	return result
}

