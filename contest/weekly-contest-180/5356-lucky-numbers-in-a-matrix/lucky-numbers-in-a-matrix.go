package leetcode

func luckyNumbers (matrix [][]int) []int {
	result := []int{}

	for i := 0 ; i < len(matrix); i ++ {
		tmpMin := matrix[i][0]
		tmpJ := 0
		for j := 1; j < len(matrix[i]); j ++ {
			if matrix[i][j] < tmpMin {
				tmpMin = matrix[i][j]
				tmpJ = j
			}
		}
		flag := true
		for z:= 0; z < len(matrix) ; z++  {
			if matrix[z][tmpJ] > tmpMin {
				flag = false
			}
		}

		if flag {
			result = append(result, tmpMin)
		}
	}
	return result
}