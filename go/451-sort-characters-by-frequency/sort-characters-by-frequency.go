package leetcode


import "sort"


func frequencySort(s string) string {

	letterMap := map[string]int{}
	for _ , value := range s {
		letterMap[string(value)] ++
	}
	var tmp  []struct {
		l     string
		count int
	}
	for key , value := range letterMap {
		tmp = append(tmp, struct {
			l     string
			count int
		}{l: key, count: value})
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i].count > tmp[j].count
	})
	result := ""
	for _ , value := range tmp {
		for i := 0 ; i < value.count; i ++ {
			result += value.l
		}
	}
	return result
}


