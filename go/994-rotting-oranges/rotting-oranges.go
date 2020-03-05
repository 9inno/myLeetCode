package leetcode

import "strconv"


func OrangesRotting(grid [][]int) int {
	rottenMap := make(map[string] []int)
	freshMap := make(map[string] []int)

	for i , row := range grid {
		for j , value := range row {
			if value == 1 {
				freshMap[strconv.Itoa(i) + "," + strconv.Itoa(j)] = [] int {i , j}
			} else if value == 2 {
				rottenMap[strconv.Itoa(i) + "," + strconv.Itoa(j)] = [] int {i, j}
			}
		}
	}
	result := 0
	for len(rottenMap) != 0 {
		tmpMap := make(map[string] []int)
		for _ , coordinates := range rottenMap{
			var ok bool
			_, ok = freshMap[strconv.Itoa(coordinates[0] + 1) + "," +strconv.Itoa(coordinates[1])]
			if ok {
				tmpMap[strconv.Itoa(coordinates[0] + 1) + "," +strconv.Itoa(coordinates[1])] = [] int { coordinates[0] + 1, coordinates[1]}
				delete(freshMap, strconv.Itoa(coordinates[0] + 1) + "," +strconv.Itoa(coordinates[1]))
			}
			_, ok = freshMap[strconv.Itoa(coordinates[0] ) + "," +strconv.Itoa(coordinates[1] + 1)]
			if ok {
				tmpMap[strconv.Itoa(coordinates[0] ) + "," +strconv.Itoa(coordinates[1]+ 1)] = [] int { coordinates[0], coordinates[1]+ 1}
				delete(freshMap, strconv.Itoa(coordinates[0] ) + "," +strconv.Itoa(coordinates[1] + 1))
			}
			_, ok = freshMap[strconv.Itoa(coordinates[0] - 1 ) + "," +strconv.Itoa(coordinates[1])]
			if ok {
				tmpMap[strconv.Itoa(coordinates[0] - 1) + "," +strconv.Itoa(coordinates[1])] = [] int { coordinates[0] - 1, coordinates[1]}
				delete(freshMap, strconv.Itoa(coordinates[0] - 1 ) + "," +strconv.Itoa(coordinates[1]))
			}
			_, ok = freshMap[strconv.Itoa(coordinates[0]  ) + "," +strconv.Itoa(coordinates[1]- 1)]
			if ok {
				tmpMap[strconv.Itoa(coordinates[0] ) + "," +strconv.Itoa(coordinates[1]- 1)] = [] int { coordinates[0], coordinates[1]- 1}
				delete(freshMap, strconv.Itoa(coordinates[0]) + "," +strconv.Itoa(coordinates[1]- 1))
			}
		}
		if len(tmpMap) != 0 {
			result = result + 1
		}
		rottenMap = tmpMap
	}

	if len(freshMap) > 0 {
		return -1
	}
	return result
}
