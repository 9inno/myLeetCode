package leetcode

func patternMatching(pattern string, value string) bool {
	countMap := map[string]int{"a":0, "b" : 0}
	for _ , v := range pattern {
		countMap[string(v)] ++
	}
	countA := countMap["a"]
	countB := countMap["b"]
	if (countA == 0 || countB == 0)  && value == "" {
		return true
	}
	if countA != 0 && countB !=0 && value == ""{
		return false
	}
	outLoop: for i := 0; i < len(value) + 1; i ++ {
		surplus := len(value) - (countA * i)

		var j int
		if countB == 0{
			j = 0
		} else {
			if surplus < 0 || surplus % countB != 0  {
				continue outLoop
			}
			j = surplus / countB
		}
		tmpA := ""
		tmpB := ""
		p := 0
		tmpStr := ""
		for _ , v := range pattern {
			if p > len(value) {
				continue outLoop
			}
			if v == 'a' {
				if p + i  > len(value) {
					continue outLoop
				}
				if tmpA == "" {
					tmpA = value[p : p + i]
				} else {
					if tmpA != value[p : p + i] {
						continue outLoop
					}
				}
				tmpStr += value[p : p + i]
				p += i
			}

			if v == 'b' {
				if p + j - 1 > len(value) {
					continue outLoop
				}
				if tmpB == "" {
					tmpB = value[p : p + j]
				}else  {
					if tmpB != value[p: p + j] {
						continue outLoop
					}
				}
				tmpStr += value[p : p + j ]
				p += j
			}
		}
		if value == tmpStr {
			return true
		}
	}

	return false
}
