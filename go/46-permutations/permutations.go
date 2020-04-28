package leetcode


func permute(nums []int) [][]int {
	
	var result [][]int
	recursion(&result, nums, 0, len(nums), []int{})
	return result
}

func recursion(result *[][]int, nums []int, n, count int, tmp []int)  {
	if count == 0 {
		return
	}
	if count == n {
		*result = append(*result, tmp)
	}

	for i := n ; i < count ; i ++ {
		nums[i] , nums[n] = nums[n] , nums[i]
		recursion(result, nums, n + 1, count, append(tmp, nums[n]))
		nums[i] , nums[n] = nums[n] , nums[i]
	}
}
