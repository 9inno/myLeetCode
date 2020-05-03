package main

func main() {
	a := []int{0,1,0,1}
	kLengthApart(a, 1)
}

func kLengthApart(nums []int, k int) bool {
	i := 0
	j := 0
	count := 0
	if nums[0] == 1 {
		j = 1
	}
	for j < len(nums) {
		if nums[j] == 1 {
			if count < k {
				return false
			} else {
				count = 0
				i = j
			}
		} else {
			count ++
		}
		j ++
	}
	if i != len(nums) -1 &&count < k {
		return false
	}
	return true
}
