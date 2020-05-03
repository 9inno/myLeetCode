package leetcode

func destCity(paths [][]string) string {
	start := map[string]bool{}
	end := map[string]bool{}
	for _ , value := range paths {
		start[string(value[0])] = true
		end[string(value[1])] = true
	}
	for key , _ := range end {
		if _, ok := start[key] ; ok {
			continue
		} else {
			return key
		}
	}
	return ""
}
