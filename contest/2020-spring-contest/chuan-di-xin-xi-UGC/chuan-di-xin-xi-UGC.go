package main

func main() {
	relation := [][]int {[]int{0,2},[]int{2,1},[]int{3,4},[]int{2,3},[]int{1,4},[]int{2,0},[]int{0,4}}
	numWays(5, relation, 3)
}

func numWays(n int, relation [][]int, k int) int {
	hashMap := map[int][]int{}
	for _ , row := range relation {
		hashMap[row[0]] = append(hashMap[row[0]], row[1])
	}
	result := 0
	recursion(hashMap, 0, n - 1, k, 0, &result)
	return result

}

func recursion(hashMap map[int][]int , index int, target int , k int, current int, result *int)  {

	if current > k {
		return
	}
	for _ , value :=range hashMap[index] {
		if k == current + 1 && target == value {
			*result ++
			return
		}
		recursion(hashMap, value, target, k , current + 1, result)
	}
}


