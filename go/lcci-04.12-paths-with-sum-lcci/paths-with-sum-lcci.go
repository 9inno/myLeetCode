package leetcode

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	var result [][]int
	recursion(root, sum, 0, []int{}, &result)
	return len(result)
}

func recursion(root *TreeNode, sum int, current int , tmp []int, result *[][]int) {
	if root == nil {
		return
	}

	if current + root.Val == sum {
		*result = append(*result, tmp)
	}
	if len(tmp) != 0 {
		recursion(root.Left, sum, current + root.Val, append(tmp, root.Val), result)
		recursion(root.Right, sum, current + root.Val, append(tmp, root.Val), result)
	}else {
		recursion(root.Left, sum, current + root.Val, append(tmp, root.Val), result)
		recursion(root.Right, sum, current + root.Val, append(tmp, root.Val), result)
		recursion(root.Left, sum, current , tmp, result)
		recursion(root.Right, sum, current , tmp, result)
	}

}
