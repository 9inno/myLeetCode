package leetcode


func Rotate(matrix [][]int)  {
	lenI := len(matrix) - 1
	if lenI == 0 {
		return
	}

	lenJ := len(matrix[0]) - 1

	i:=0
	j:=0

	for {
		for z := 0 ; z < lenJ - j; z++ {
			tmp := matrix[i][j + z]
			matrix[i][j + z] = matrix[lenI - z][j]
			matrix[lenI - z][j] = matrix[lenI][lenJ - z]
			matrix[lenI][lenJ - z] = matrix[i + z][lenJ]
			matrix[i + z][lenJ] = tmp
		}
		i ++
		j ++
		lenI --
		lenJ --
		if i >= lenI {
			break
		}
	}
}

//找四个顶点  顺时针替换 每条边同时 顺时针位置 + 1  一圈一圈遍历到中心停止
