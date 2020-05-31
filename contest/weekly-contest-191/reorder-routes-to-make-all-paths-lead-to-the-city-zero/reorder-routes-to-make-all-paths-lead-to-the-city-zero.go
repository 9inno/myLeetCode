package main

func main() {
//	[[0,1],[1,3],[2,3],[4,0],[4,5]]

	a := [][]int{ []int{0,1}, []int{1,3}, []int{2,3}, []int{4,0}, []int{4,5}}
	minReorder(5, a)
}


func minReorder(n int, connections [][]int) int {
	target := map[int]int {0: 0}
	result := 0
	count := n - 1
	for count > 0{
		for _ , value := range connections{
			if _ , ok := target[value[0]]; ok {
				result ++
				count --
				target[value[1]] = 0
			} else if _ , ok := target[value[1]]; ok {
				count --
				target[value[0]] = 0
			}
		}
	}
	return result
}
