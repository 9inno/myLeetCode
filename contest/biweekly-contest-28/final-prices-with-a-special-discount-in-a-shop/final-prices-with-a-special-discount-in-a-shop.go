package leetcode

func finalPrices(prices []int) []int {
	result := make([]int , 0)
	for i:= 0; i < len(prices); i ++ {
		tmp := 0
		for j := i + 1; j < len(prices); j ++ {
			if prices[i] >= prices[j] {
				tmp = prices[j]
				break
			}
		}
		result = append(result, prices[i] - tmp)
	}
	return result
}
