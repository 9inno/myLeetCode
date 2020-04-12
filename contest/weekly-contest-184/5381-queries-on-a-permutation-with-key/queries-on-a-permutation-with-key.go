package main

func main() {
	processQueries([]int{3,1,2,1}, 5)
}

func processQueries(queries []int, m int) []int {
	var p []int
	var result []int
	for i := 1; i < m + 1; i ++ {
		p = append(p, i)
	}

	for _ , value :=range queries {

		for j := 0 ; j < len(p); j ++ {
			tmp := p[j]
			if tmp == value {
				result = append(result, j)
				p = append([]int{tmp}, append(p[0:j], p[j + 1:]...)...,)
			}
		}
	}

	return result
}
