package leetcode



type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		tmpQueue := []*TreeNode{}
		for _ , value := range queue {
			if value != nil {
				tmpQueue = append(tmpQueue, value.Left)
				tmpQueue = append(tmpQueue, value.Right)
			}
		}
		i, j := 0, len(tmpQueue) - 1
		for i < j {
			if tmpQueue[i] == nil || tmpQueue[j] == nil {
				if tmpQueue[i] != tmpQueue[j] {
					return false
				}
			}else {
				if tmpQueue[i].Val != tmpQueue[j].Val {
					return false
				}
			}
			i ++
			j --
		}
		queue = tmpQueue
	}
	return true
}


func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}
