package main

func main() {
	a := []int{-3,2,-3,4,2}
	minStartValue(a)
}

func minStartValue(nums []int) int {

	i := 1
	outLoop :for {
		sum := i
		for _ , value := range nums {
			sum += value
			if sum < 1 {
				i++
				continue outLoop
			}
		}
		break
	}
	return i
}
