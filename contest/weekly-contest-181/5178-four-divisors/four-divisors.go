package main

import "fmt"

func main() {
	a:= sumFourDivisors([]int{90779,36358,90351,75474,32986})
	fmt.Print(a)
}
func sumFourDivisors(nums []int) int {
	result := 0
	for _ , value := range nums{
		count := 0
		tmpSum := 0
		for i := 1; i *i <= value; i++ {
			if value % i == 0 {
				count++
				tmpSum+=i
				if i*i < value {
					count++
					tmpSum += value/i
				}
			}
		}
		if count == 4 {
			 result += tmpSum
		}
	}
	return result
}