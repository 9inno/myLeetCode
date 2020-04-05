package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(numSteps("10"))
}
func numSteps(s string) int {
	result, _ :=strconv.ParseInt(s,2, 64)

	return numberOfSteps(int(result)) + 1
}

func numberOfSteps(num int) int {
	if num > 1 {
		if num % 2 == 1 {
			return  1 + numberOfSteps( (num + 1)>>1 )
		} else {
			return  1 + numberOfSteps( num>>1 )
		}
	}
	return num
}
