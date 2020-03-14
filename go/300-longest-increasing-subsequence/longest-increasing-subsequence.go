package leetcode

func LengthOfLIS(nums []int) int {
	if len(nums)==0 {
		return 0
	}

	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]
	length := 1

	for k:=1; k<len(nums); k++ {
		if nums[k]>dp[length] {
			length++
			dp[length] = nums[k]
		}else {
			// 二分查找
			l := 1
			r := length
			for l<r {
				m := l + (r-l)/2
				if dp[m]>=nums[k] {
					r = m
				}else{
					l = m+1
				}
			}

			dp[l] = nums[k]
		}
	}
	return length
}