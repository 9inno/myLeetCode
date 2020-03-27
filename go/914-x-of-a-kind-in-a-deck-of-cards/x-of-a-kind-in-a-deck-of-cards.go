package leetcode

func HasGroupsSizeX(deck []int) bool {
	hashMap := make(map[int]int)
	for _ , value := range deck {
		if _, ok := hashMap[value]; ok{
			hashMap[value] ++
		} else {
			hashMap[value] = 1
		}
	}
	var slice []int
	for _  , count := range hashMap {
		slice = append(slice, count)
	}

	return sliceGcd(slice, len(slice)) >= 2
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a % b)
	}
}

func sliceGcd(slice []int, len int) int {
	if len == 1 {
		return slice[0]
	}
	return gcd(slice[len -1], sliceGcd(slice , len -1))
}
