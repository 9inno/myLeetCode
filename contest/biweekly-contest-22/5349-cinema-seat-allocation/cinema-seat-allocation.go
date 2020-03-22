package leetcode


func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	hashMap := map[int] []int{}
	result := 0
	for _ , value := range reservedSeats{
		hashMap[value[0]] = append(hashMap[value[0]], value[1])
	}
	tmpCount := 0
	for _ , row := range hashMap {
		var tmpSlice [11]int
		for _ , column := range row {
			tmpSlice[column] = 1
		}
		if (tmpSlice[2] == 1 || tmpSlice[3] == 1 || tmpSlice[8] == 1 || tmpSlice[9] == 1) && (tmpSlice[4] != 1 && tmpSlice[5] != 1 && tmpSlice[6] != 1 && tmpSlice[7] != 1) {
			result = result + 1
		}
		if tmpSlice[2] != 1 && tmpSlice[3] != 1 && tmpSlice[4] != 1 && tmpSlice[5] != 1 && tmpSlice[6] != 1 && tmpSlice[7] != 1 && tmpSlice[8] != 1 && tmpSlice[9] != 1{
			result = result + 2
		}
		if ((tmpSlice[5] == 1 || tmpSlice[4] == 1)  && tmpSlice[6] != 1 && tmpSlice[7] != 1 && tmpSlice[8] != 1 && tmpSlice[9] != 1 ) || ((tmpSlice[6] == 1 || tmpSlice[7] == 1) && tmpSlice[2] != 1 && tmpSlice[3] != 1 && tmpSlice[4] != 1 && tmpSlice[5] != 1 ){
			result = result + 1
		}
		tmpCount ++
	}
	result = result + (n - tmpCount ) * 2
	return result
}