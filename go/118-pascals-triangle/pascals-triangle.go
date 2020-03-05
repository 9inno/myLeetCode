package leetcode


func Generate(numRows int) [][]int {
	if numRows == 0 {
		slice := make([][]int, 0)
		return slice
	}
	if numRows == 1 {
		slice := [][]int{[]int{1}}
		return slice
	}

	tmp := Generate(numRows - 1)
	tmpSlice := make([]int, numRows)
	for i := 0; i < numRows ; i = i + 1 {
		if i == 0 || i == (numRows - 1) {
			tmpSlice[i] = 1
		} else {
			tmpSlice[i] = tmp[numRows - 2][i - 1] + tmp[numRows - 2][i]
		}
	}

	return append(tmp,tmpSlice)

}