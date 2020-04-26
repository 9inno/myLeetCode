package leetcode


func findDiagonalOrder(nums [][]int) []int {
	n := len(nums)

	if n == 0  {
		return []int{}
	}
	type s struct {
		x int
		y int
	}
	var result []int
	queue := []s{{x: 0, y: 0}}
	hashMap := map[s]bool{}
	for len(queue) != 0 {
		var tmpQueue []s

		for _ , value := range queue {
			result = append(result , nums[value.x][value.y])
			if value.x + 1 < n - 1 && value.y < len(nums[value.x + 1]) {
				tmpStruct :=struct {
					x int
					y int
				}{x: value.x + 1, y: value.y}
				if _, ok := hashMap[tmpStruct] ; !ok {
					hashMap[tmpStruct] = true
					tmpQueue = append(tmpQueue, tmpStruct)
				}

			}
			if value.y + 1 < len(nums[value.x]) {
				tmpStruct :=struct {
					x int
					y int
				}{x: value.x , y: value.y + 1}
				if _, ok := hashMap[tmpStruct] ; !ok {
					hashMap[tmpStruct] = true
					tmpQueue = append(tmpQueue, tmpStruct)
				}
			}
		}
		queue = tmpQueue
	}
	return result
}
