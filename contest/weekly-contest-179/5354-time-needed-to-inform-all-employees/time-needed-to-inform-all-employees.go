package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	result := 0
	for i := 0; i < len(manager); i = i + 1 {
		if informTime[i] == 0 {
			tmp := 0
			tmpI := i
			for tmpI != -1  {
				tmp = tmp + informTime[tmpI]
				tmpI = manager[tmpI]
			}
			if result < tmp {
				result = tmp
			}
		}
	}
	return result
}
