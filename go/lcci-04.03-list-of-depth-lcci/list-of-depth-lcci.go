package leetcode


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

type ListNode struct {
    Val int
    Next *ListNode
}

func listOfDepth(tree *TreeNode) []*ListNode {
	var result []*ListNode
	queue := []*TreeNode{tree}

	for len(queue) != 0 {
		var tmpQueue []*TreeNode
		list := &ListNode{
			Val:  0,
			Next: nil,
		}
		tmpList := list
		for _ , node := range queue {
			if node != nil {
				list.Next = &ListNode{
					Val:  node.Val,
					Next: nil,
				}
				list = list.Next
				if node.Left != nil {
					tmpQueue = append(tmpQueue, node.Left)
				}
				if node.Right != nil {
					tmpQueue = append(tmpQueue, node.Right)
				}
			}
		}
		queue = tmpQueue
		if tmpList.Next != nil {
			result = append(result, tmpList.Next)
		}
	}
	return result
}
