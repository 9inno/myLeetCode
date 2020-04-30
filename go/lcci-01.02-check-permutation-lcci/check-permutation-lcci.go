package leetcode


func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	hashMapS1 := map[int32]int{}
	hashMapS2 := map[int32]int{}

	for _ , value := range s1 {
		hashMapS1[value]++
	}
	for _ , value := range s2 {
		hashMapS2[value]++
	}

	for key, value := range hashMapS1 {
		if tmp, ok := hashMapS2[key] ; ok {
			if tmp != value{
				return false
			}
		}else {
			return false
		}
	}
	return true
}

