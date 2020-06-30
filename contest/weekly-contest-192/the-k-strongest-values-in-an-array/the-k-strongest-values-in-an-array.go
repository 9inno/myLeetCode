package main

import (
	"math"
	"sort"
)

func main() {
	a := []int{6,7,11,7,6,8}
	getStrongest(a, 4)
}

func getStrongest(arr []int, k int) []int {
	sort.Ints(arr)
	m := arr[(len(arr) - 1) /2]

	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i] - m )) > math.Abs(float64(arr[j] - m)) || (math.Abs(float64(arr[i] - m )) == math.Abs(float64(arr[j] - m)) &&arr[i] > arr[j])
	})

	return arr[: k]
}
