package leetcode

func setZeroes(matrix [][]int)  {
	var queue []struct{ x, y int }

	for i , row := range matrix {
		for j , column := range row {
			if column == 0 {
				queue = append(queue, struct{ x, y int }{x: i, y: j})
			}
		}
	}

	for _ , value := range queue {
		for z := 0; z < len(matrix[value.x]) ; z ++ {
			if  matrix[value.x][z] == 0{
				continue
			}
			matrix[value.x][z] = 0
		}
		for z:= 0 ; z < len(matrix) ; z ++ {
			if  matrix[z][value.y] == 0{
				continue
			}
			matrix[z][value.y] = 0
		}
	}

}
