package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func pseudoPalindromicPaths (root *TreeNode) int {
	result := 0
	hashMap := map[int]int{}
	recursion(root, hashMap, &result)
	return result
}

func recursion(root *TreeNode, hashMap map[int]int, result *int)  {
	if root == nil {
		return
	}
	hashMap[root.Val] ++
	if root.Left == nil && root.Right == nil {
		tmpValue := 0
		for _ , value := range hashMap {
			if value % 2 != 0 {
				tmpValue ++
			}
		}
		if tmpValue < 2 {
			*result++
		}
		return
	}
	tmpLeftMap := map[int]int{}
	tmpRightMap := map[int]int{}
	for key , value := range hashMap {
		tmpLeftMap[key] = value
		tmpRightMap[key] = value
	}

	recursion(root.Left, tmpLeftMap, result)
	recursion(root.Right, tmpRightMap, result)
}
