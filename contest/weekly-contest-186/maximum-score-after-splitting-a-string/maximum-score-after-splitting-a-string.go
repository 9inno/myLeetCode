package leetcode

func maxScore(s string) int {
	right := 0
	left := 0
	for _ , value := range s {
		if string(value) == "1" {
			right ++
		}
	}
	max := 0
	for i := 0 ; i <len(s) - 1; i ++ {
		if string(s[i]) == "0" {
			left ++
		}else  {
			right --
		}

		max = Max(max, left + right)
	}


	return max

}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
