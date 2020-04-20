package leetcode



type s struct {
	x int
	y int
}
func NumIslands(grid [][]byte) int {

	result := 0
	hashMap := map[s]bool{}

	for i , row :=range grid {
		for j , column := range row {
			tmpStruct := s{
				x: i,
				y: j,
			}
			if _, ok := hashMap[tmpStruct]; ok {
				continue
			} else {
				hashMap[tmpStruct] = true
			}

			if column == 49 {
				queue := []s{tmpStruct}
				result ++
				for len(queue) != 0 {
					tmpQueue := []s{}

					for _ , value := range queue {
						//上
						if value.x > 0 {
							element := s{
								x: value.x - 1,
								y: value.y,
							}
							if _, ok := hashMap[element] ; !ok && grid[element.x][element.y] == 49 {
								tmpQueue = append(tmpQueue, element)
								hashMap[element] = true
							}
						}
						// 下
						if value.x < len(grid) - 1 {
							element := s{
								x: value.x  + 1,
								y: value.y ,
							}
							if _, ok := hashMap[element] ; !ok && grid[element.x][element.y] == 49 {
								tmpQueue = append(tmpQueue, element)
								hashMap[element] = true
							}
						}
						//左
						if value.y > 0 {
							element := s{
								x: value.x,
								y: value.y - 1,
							}
							if _, ok := hashMap[element] ; !ok && grid[element.x][element.y] == 49 {
								tmpQueue = append(tmpQueue, element)
								hashMap[element] = true
							}
						}
						//右
						if value.y < len(row) - 1 {
							element := s{
								x: value.x,
								y: value.y + 1,
							}
							if _, ok := hashMap[element] ; !ok && grid[element.x][element.y] == 49 {
								tmpQueue = append(tmpQueue, element)
								hashMap[element] = true
							}
						}
					}

					queue = tmpQueue
				}
			}
		}
	}

	return result
}



func NumIslands1(grid [][]byte) int {

	result := 0
	for i , row :=range grid {
		for j , column := range row {
			tmpStruct := s{
				x: i,
				y: j,
			}
			if column == 49 {
				result ++
				recursion(tmpStruct, grid)
			}

		}
	}
	return result
}
func recursion( element s, grid [][]byte) {

	if 0 <= element.x && element.x < len(grid) && 0 <= element.y && element.y < len(grid[0]) && grid[element.x][element.y] == 49 {
		grid[element.x][element.y] = 48
		recursion(s{x: element.x - 1, y: element.y},grid)
		recursion(s{x: element.x + 1, y: element.y},grid)
		recursion(s{x: element.x , y: element.y- 1},grid)
		recursion(s{x: element.x , y: element.y + 1},grid)
	}
}