package leetcode


func getNoZeroIntegers(n int) []int {
	i := 1
	for  {
		if !hasZero(i) && !hasZero(n - i) {
			break
		}
		i ++
	}
	return []int{i , n - i}
}

func hasZero(num int) bool {
	for num != 0{
		if num % 10 == 0 {
			return true
		}
		num /= 10
	}
	return false
}
