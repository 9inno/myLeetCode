package leetcode

func dailyTemperatures(T []int) []int {
	stack := []struct{
		key int
		v int
	}{}
	result := make([]int, len(T))
	for key , value := range T {
		if len(stack) != 0 && stack[len(stack) - 1].v < value {
			for len(stack) != 0 || stack[len(stack) - 1].v < value {
				result[stack[len(stack) - 1].key] = key - stack[len(stack) - 1].key
				stack = stack[:len(stack) - 1]
			}
		}
		stack = append(stack, struct {
			key int
			v   int
		}{key: key, v: value})
	}

	for _ , value := range stack {
		result[value.key] = 0
	}
	return result
}
