package leetcode

import "sort"

func wiggleSort(nums []int)  {
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	var tmp []int
	for i > j {
		tmp = append(tmp, nums[i])
		tmp = append(tmp, nums[j])
		i++
		j--
	}
	copy(nums, tmp)
}


func wiggleSort1(nums []int)  {
	for i := 0; i < len(nums)-1; i++ {
		if (i%2 == 1 && nums[i] < nums[i+1]) || (i%2 == 0 && nums[i] > nums[i+1]) {
			nums[i], nums[i+1] = nums[i+1], nums[i]
		}
	}
}