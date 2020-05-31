package leetcode

func SubarraysDivByK(A []int, K int) int {
	hashMap := map[int]int {0 : 1}
	pre := 0
	result := 0

	for _ , value := range A {
		pre += value

		mod := (pre % K + K) % K
		if count , ok := hashMap[mod]; ok {
			result += count
		}
		hashMap[mod] ++
	}
	return result
}
