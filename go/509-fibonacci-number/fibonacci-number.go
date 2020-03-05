package leetcode

func Fib(N int) int {

	hashMap := make(map[int] int)
	type fc func(index int, tmpMap map[int]int) int
	var tmpFunc fc
	tmpFunc = func(index int, tmpMap map[int]int) int {
		if result , ok := tmpMap[index] ; ok {
			return result
		}

		if index < 2 {
			return index
		}

		result := tmpFunc(index - 1, tmpMap) + tmpFunc(index - 2, tmpMap)
		tmpMap[index] = result
		return result
	}
	return tmpFunc(N, hashMap)

}