package leetcode

func majorityElement(nums []int) int {
	hashMap := map[int]int{}
	len := len(nums) / 2
	for _ , value := range nums {
		hashMap[value] ++
	}
	for key , count := range hashMap {
		if count > len {
			return key
		}
	}
	return - 1
}