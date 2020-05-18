package main

func main() {
	a := []int{1,1,1}
	subarraySum(a, 2)
}


func subarraySum(nums []int, k int) int {

	result := 0
	for i := 0; i < len(nums); i ++ {
		tmpSum := 0
		for j := i; j <len(nums); j ++ {
			tmpSum += nums[j]
			if tmpSum == k {
				result ++
			}
		}
	}
	return result
}
