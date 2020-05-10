package main

func main() {
	countTriplets([]int{2,3,1,6,7})
}

func countTriplets(arr []int) int {
	result := 0
	for i := 0; i < len(arr); i++ {
		tmp := arr[i]
		for j := i + 1; j < len(arr); j++ {
			tmp ^= arr[j]
			if tmp == 0 {
				result += j - i
			}

		}
	}
	return result
}