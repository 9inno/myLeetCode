package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Print(minSubsequence([]int{4,3,10,9,8}))
	fmt.Print(minSubsequence([]int{6}))

}

func minSubsequence(nums []int) []int {
	sum := 0
	for _ , value := range nums {
		sum += value
	}
	sort.Ints(nums)
	var result []int
	target := sum /2
	for i:= len(nums) - 1; i >= 0 ; i -- {
		if target >= 0 {
			result = append(result, nums[i])
			target -= nums[i]
		} else  {
			break
		}
	}
	return  result
}







