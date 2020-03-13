package leetcode

import "sort"

func MajorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums) / 2]

}


func MajorityElement1(nums []int) int {
	hashMap := make(map[int] int)

	for _ , value := range nums{
		if _ , ok := hashMap[value] ; ok {
			hashMap[value] ++
		} else {
			hashMap[value] = 1
		}
		if hashMap[value] > len(nums) / 2 {
			return value
		}
	}
	return 0
}