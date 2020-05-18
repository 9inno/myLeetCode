package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(gcd(1,6))
}
func simplifiedFractions(n int) []string {
	result := []string{}

	for i := 1; i < n; i ++ {
		for j := i + 1; j <= n ; j ++ {
			if gcd(i, j) == 1 {
				result = append(result, strconv.Itoa(i) + "/" + strconv.Itoa(j))
			}
		}
	}
	return result
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a % b)
	}
}
