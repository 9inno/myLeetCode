package main

func main() {
	increase := [][]int {
		[]int{0,4,5},
		[]int{4,8,8},
		[]int{8,6,1},
		[]int{10,10,0},
	}

	requirements := [][]int {
		[]int{12,11,16},
		[]int{20,2,6},
		[]int{9,2,6},
		[]int{10,18,3},
		[]int{8,14,9},
	}
	getTriggerTime(increase, requirements)
}

func getTriggerTime(increase [][]int, requirements [][]int) []int {
	increase = append([][]int{[]int{0,0,0}}, increase...)
	for i := 1 ; i < len(increase); i ++ {
		increase[i][0] += increase[i - 1][0]
		increase[i][1] += increase[i - 1][1]
		increase[i][2] += increase[i - 1][2]
	}
	var result []int
	for _ , input := range requirements {
		result = append(result, find(0 , len(increase) - 1, input, increase))
	}
	return result
}

func compare(a, b []int) bool {
	for i := 0; i < len(a); i ++ {
		if a[i] < b[i] {
			return false
		}
	}
	return true
}

func find(start, end int , input []int, increase [][]int) int {
	if end-start == 1 {
		if compare(increase[start], input) {
			return start
		} else if compare(increase[end], input) {
			return end
		}
		return -1
	}
	if compare(increase[(start + end)/2], input) {
		return find(start, (start+end) /2, input, increase)
	}else  {
		return find((start+end) /2, end, input, increase)
	}
}