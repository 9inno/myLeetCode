package leetcode


func IsUnique(astr string) bool {

	hashMap := map[int32]bool{}

	for _ , value := range astr {
		if _, ok := hashMap[value]; ok {
			return false
		}else {
			hashMap[value] = true
		}
	}
	return true
}