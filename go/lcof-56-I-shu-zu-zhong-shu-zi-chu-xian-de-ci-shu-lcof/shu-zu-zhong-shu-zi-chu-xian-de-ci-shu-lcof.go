package leetcode

func singleNumbers(nums []int) []int {
	tmp := 0
	a := 0
	b := 0
	for _ , value := range nums {
		tmp ^= value
	}
	n := 1
	for tmp & n == 0 {
		n <<= 1
	}
	for _ , value :=range nums {
		if n & value == 0 {
			a ^= value
		} else {
			b ^= value
		}
	}
	return []int{a, b}
}
