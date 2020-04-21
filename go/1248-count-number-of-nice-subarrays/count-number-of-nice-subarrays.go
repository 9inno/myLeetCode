package leetcode


//超时 
func NumberOfSubarrays(nums []int, k int) int {
	result := 0
	recursion(nums, k , false, 0 , &result)
	return result
}

func recursion(nums []int, k int, flag bool,count int, result *int) {

	if count == k {
		*result ++
	} else if count > k {
		return
	}

	if len(nums) > 0 {
		if flag {
			recursion(nums[1:], k , flag, count + (nums[0] %2), result)
		} else {
			recursion(nums[1:], k , flag, count, result)
			recursion(nums[1:], k , !flag, count + (nums[0] %2), result)
		}
	}
}

func NumberOfSubarrays1(nums []int, k int) int {

	i := 0
	j := 0
	count := 0
	result := 0
	lastJ:= 0
	for   i < len(nums) {

		if  j < len(nums) && count + nums[j] % 2 <= k{
			count += nums[j] % 2
			if nums[j] % 2 == 1{
				lastJ = j + 1
			}
			j ++
		} else {
			count -= nums[i] % 2
			i ++
			j = lastJ
		}
		if k == count {
			result ++
		}
	}
	return result

}