package leetcode

func missingNumber(arr []int) int {

	tmp := (arr[len(arr) - 1] - arr[0] ) / len(arr)
	result := arr[0]
	for _ , value := range arr {
		if result != value {
			break
		}
		result += tmp
	}
	return result

}
