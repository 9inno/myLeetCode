package leetcode


func MyAtoi(str string) int {
	int32Max := 1 << 31 - 1
	int32Min := -1 << 31
	strIntMap := map[string]int{ "0": 0, "1": 1, "2" : 2, "3": 3, "4": 4, "5" : 5, "6" : 6, "7" : 7, "8" :8, "9" : 9 }
	sign := 1
	result := 0
	begin := false

	for _ , letter := range str {
		if value, ok := strIntMap[string(letter)]; ok {
			if result > int32Max / 10 || result == int32Max / 10 && value >7 {
				return int32Max
			}
			if result < int32Min /10 || result == int32Min / 10 && value > 8 {
				 return int32Min
			}
			result = result * 10 + sign * value
			begin = true
		} else {
			if begin {
				break
			} else {
				if string(letter) == "+" {
					 begin = true
				}else if string(letter) == "-" {
					begin = true
					sign = -1
				}else if string(letter) != " " {
					return 0
				}
			}
		}
	}
	return result

}