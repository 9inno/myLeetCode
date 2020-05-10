package main

func main() {
	buildArray([]int{1,3}, 3)
}

func buildArray(target []int, n int) []string {
	j := 0
	result := []string{}
	tmp := []int{}
	for i := 1 ; len(tmp) < len(target) ; i ++ {
		if j < len(target)  && target[j] == i {
			result = append(result, "Push")
			tmp = append(tmp, i)
			j ++
		} else {
			result = append(result, "Push")
			result = append(result, "Pop")
		}
	}
	return result
}
