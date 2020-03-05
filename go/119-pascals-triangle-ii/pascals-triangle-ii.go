package leetcode

func GetRow(rowIndex int) []int {
	type fc func(index int) []int
	var tmpFunc fc
	tmpFunc = func(index int) []int {
		if index == 1 {
			return []int {1}
		}

		tmp := tmpFunc(index - 1)
		result := make([]int, index)
		for i := 0; i < index; i = i + 1 {
			if i == 0 || i == index-1 {
				result[i] = 1
			} else  {
				result[i] = tmp[i] + tmp[i - 1]
			}
		}
		return result
	}

	return tmpFunc(rowIndex + 1)

}