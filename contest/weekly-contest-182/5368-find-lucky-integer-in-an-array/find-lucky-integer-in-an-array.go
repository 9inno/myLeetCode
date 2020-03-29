package leetcode

func findLucky(arr []int) int {
	hashMap := make(map[int]int)
	for _ , value := range arr{
		if _ , ok := hashMap[value]; ok{
			hashMap[value] ++
		}else {
			hashMap[value] = 1
		}
	}
	result := -1
	max := 0
	for key , count := range hashMap {
		if key == count {
			if key > max {
				max = key
				result = count
			}
		}
	}
	return result

}
