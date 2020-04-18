package main

func minCount(coins []int) int {
	result := 0
	for _ , value := range coins {
		if value % 2 == 1 {
			result  = result + value/2 + 1
		}else {
			result = result + value / 2
		}
	}
	return result
}

