package leetcode


func maxDepthAfterSplit(seq string) []int {
	var result []int
	deep := 0
	for  _ , letter := range seq {
		if string(letter) == "(" {
			deep ++
			result = append(result, deep % 2)
		}
		if string(letter) == ")" {
			result = append(result, deep % 2)
			deep --
		}
	}
	return result
}