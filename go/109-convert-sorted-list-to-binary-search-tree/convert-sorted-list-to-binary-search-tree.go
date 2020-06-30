package leetcode


type ListNode struct {
    Val int
    Next *ListNode
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	tmp := []int{}
	for head != nil {
		tmp = append(tmp, head.Val)
		head = head.Next
	}
	return recursion(tmp, 0, len(tmp) - 1)
}

func recursion (nums []int, start , end int) *TreeNode{
	if start > end {
		return nil
	}
	mid := start + (end - start) / 2
	tree := &TreeNode{Val:nums[mid]}
	tree.Left = recursion(nums, start , mid - 1)
	tree.Right = recursion(nums, mid + 1, end)
	return tree
}