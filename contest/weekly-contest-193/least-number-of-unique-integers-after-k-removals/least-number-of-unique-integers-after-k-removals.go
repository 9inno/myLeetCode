package main


import "sort"

func main() {
	a :=[]int{5,5,4}
	findLeastNumOfUniqueInts(a , 1)
}
func findLeastNumOfUniqueInts(arr []int, k int) int {
	hashMap := map[int]int{}
	for _ , value := range arr {
		hashMap[value] ++
	}
	tmp := []int{}
	for _ , value := range hashMap {
		tmp = append(tmp, value)
	}
	sort.Ints(tmp)
	i := 0
	for  ; i < len(tmp); i ++ {
		k -= tmp[i]
		if k < 0 {
			break
		}
	}
	return len(tmp) - i
}
