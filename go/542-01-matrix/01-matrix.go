package leetcode


func updateMatrix(matrix [][]int) [][]int {
	type coordinate  struct{ x , y int}
	maxHeight := len(matrix) - 1
	if maxHeight < 0  {
		return matrix
	}
	
	maxLength := len(matrix[0]) - 1
	 for i , row := range matrix {
		loop : for j , column := range row {
			if column == 1 {
				root := coordinate{
					x: i,
					y: j,
				}
				queue := []coordinate{root}
				tmpMap := map[coordinate]bool{root:true}
				distance := 1
				for len(tmpMap) != 0 {
					var tmpQueue []coordinate
					for _ , value := range queue {
						
	
						//左
						if value.y > 0 {
							left := coordinate{
								x: value.x,
								y: value.y - 1,
							}
							if matrix[left.x][left.y] == 0 {
								matrix[i][j] = distance
								continue loop
							}
							if _, ok := tmpMap[left]; !ok {
								tmpMap[left] = true
								tmpQueue = append(tmpQueue, left) 
							}
						}
						//右			
						if value.y < maxLength {
							right := coordinate{
								x: value.x,
								y: value.y + 1,
							}
							if matrix[right.x][right.y] == 0 {
								matrix[i][j] = distance
								continue loop
							}
							if _, ok := tmpMap[right]; !ok {
								tmpMap[right] = true
								tmpQueue = append(tmpQueue, right)
							}
						}
						//上
						if value.x > 0 {
							top := coordinate{
								x: value.x - 1,
								y: value.y,
							}
							if matrix[top.x][top.y] == 0 {
								matrix[i][j] = distance
								continue loop
							}
							if _, ok := tmpMap[top]; !ok {
								tmpMap[top] = true
								tmpQueue = append(tmpQueue, top)
							}
						}
						//下
						if value.x < maxHeight {
							bottom := coordinate{
								x: value.x + 1,
								y: value.y,
							}
							if matrix[bottom.x][bottom.y] == 0 {
								matrix[i][j] = distance
								continue loop
							}
							if _, ok := tmpMap[bottom]; !ok {
								tmpMap[bottom] = true
								tmpQueue = append(tmpQueue, bottom)
							}
						}
	
					}
					distance ++
					queue = tmpQueue
				}
			}
		}
	}
	return matrix
}