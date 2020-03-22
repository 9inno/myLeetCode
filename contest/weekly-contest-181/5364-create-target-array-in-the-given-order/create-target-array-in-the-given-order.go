package main

func main() {
	createTargetArray([]int{0,1,2,3,4}, []int{0,1,2,2,1})
}
func createTargetArray(nums []int, index []int) []int {
	var result []int
	for i := 0 ; i < len(nums); i ++ {
		tmp := make([]int, len(result[index[i]:]))
		copy(tmp,result[index[i]:])
		result = append(result[0:index[i]], nums[i])
		for _, value := range tmp {
			result = append(result, value)
		}
	}
	return result
}