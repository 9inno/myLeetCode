package leetcode


func CanJump(nums []int) bool {
	k := 0
	for key , value := range nums {
		if key > k {
			return false
		}
		k = max(k, key + value)
	}
	return true
}

func CanJump1(nums []int) bool {
	k := 0
	for i := 0; k < len(nums) - 1; i ++ {
		if i > k {
			return false
		}
		k = max(k , i + nums[i])
	}
	return true
}

func max(a , b int) int  {
	if a > b {
		return a
	}
	return b
}