package leetcode

func FindContinuousSequence(target int) [][]int {
	left := 1
	right := 1
	sum := 0
	var result  [][]int
	for left <= target / 2  {
		if sum < target {
			sum = sum + right
			right = right + 1
		}else if sum > target {
			sum = sum - left
			left = left + 1
		} else {
			var tmp []int
			for i := left; i < right; i = i + 1 {
				tmp = append(tmp , i)
			}
			result = append(result ,tmp)
			sum = sum - left
			left = left + 1
		}
	}

	return result
}