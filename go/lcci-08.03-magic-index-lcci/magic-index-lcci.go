package leetcode

func findMagicIndex(nums []int) int {
	for key , value := range nums {
		if key == value {
			return key
		}
	}
	return -1
}
