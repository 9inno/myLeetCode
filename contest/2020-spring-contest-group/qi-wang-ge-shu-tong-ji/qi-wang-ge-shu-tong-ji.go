package main

func main() {
	a:= []int{1,2,3,3,3,5,5,5,5}
	expectNumber(a)
}
func expectNumber(scores []int) int {

	hashMap := map[int]int{}
	for _ , value := range scores {
		hashMap[value] ++
	}

	return len(hashMap)
}
