package main

func main() {
	a := [][]int{[]int{1,2}, []int{1, 0}, []int{2,0}}
	q := [][]int{[]int{1,0}, []int{1,2}}
	checkIfPrerequisite(3, a, q)
}

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	hashMap := map[int][]int{}

	for _ , value := range prerequisites {
		if _ , ok := hashMap[value[0]] ; ok {
			hashMap[value[0]] = append( hashMap[value[0]], value[1])
		} else {
			hashMap[value[0]] = []int{value[1]}
		}
	}
	var result []bool

	for _ ,value := range queries{
		if _ , ok := hashMap[value[0]]; ok {
			result = append(result, check(hashMap, [][]int{hashMap[value[0]]}, value[1]))
		} else {
			result = append(result, false)
		}
	}

	return result
}

func check(hashMap map[int][]int, queue [][]int, target int) bool {

	tmpMap := map[int]bool{}

	for len(queue) != 0 {
		tmpQueue := [][]int{}
		for _ ,value := range queue {
			for _ , v := range value {
				if v == target {
					return true
				}else {
					if _, ok := tmpMap[v] ; !ok {
						tmpQueue = append(tmpQueue, hashMap[v])
						tmpMap[v] = true
					}
				}
			}
		}
		queue = tmpQueue
	}
	return false
}
