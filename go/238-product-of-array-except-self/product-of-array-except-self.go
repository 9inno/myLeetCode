package main

func main() {
	a := []int {1,2,3,4}
	productExceptSelf1(a)
}
func productExceptSelf(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	left := make([]int , length)
	right := make([]int , length)
	result := make([]int , length)
	left[0] = 1
	right[length- 1] = 1
	for i := 1; i <len(nums) ; i ++ {
		left[i] = left[i - 1] * nums[i - 1]
	}
	result[length - 1] = left[length - 1]
	for i := length - 2 ; i >= 0; i -- {
		right[i] = right[i + 1] * nums[i + 1]
		result[i] = right[i] * left[i]
	}

	return result
}

func productExceptSelf1(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	left := 1
	right := 1
	result := make([]int , length)
	result[0] , result[length - 1] = 1 , 1

	for i := 1; i <len(nums) ; i ++ {
		left *= nums[i - 1]
		result[i] = left
	}
	for i := length - 2 ; i >= 0; i -- {
		right *= nums[i + 1]
		result[i] *= right
	}

	return result
}