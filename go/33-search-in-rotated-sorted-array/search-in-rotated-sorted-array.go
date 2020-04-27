package leetcode



func Search(nums []int, target int) int {

	return recursion(&nums, 0, len(nums) -1 ,target)
}

func recursion(nums *[]int, start, end ,target int) int {
	if start > end {
		return -1
	}
	mid := start + (end -start)/ 2

	if  (*nums)[start] == target{
		return start
	}
	if  (*nums)[end] == target{
		return end
	}
	if  (*nums)[mid] == target{
		return mid
	}

	if (*nums)[start] <= (*nums)[mid] {
		if (*nums)[start] < target && (*nums)[mid] > target {
			return recursion(nums , start + 1, mid -1 , target)
		} else {
			return recursion(nums, mid + 1, end - 1, target)
		}
	} else {
		if (*nums)[mid] < target && (*nums)[end] > target {
			return recursion(nums , mid + 1, end -1 , target)
		} else {
			return recursion(nums, start + 1, mid - 1, target)
		}
	}
}