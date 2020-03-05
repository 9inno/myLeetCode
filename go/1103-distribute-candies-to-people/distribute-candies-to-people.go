package leetcode


func DistributeCandies(candies int, numPeople int) []int {

	result := make([]int, numPeople)
	i := 1
	j := 0
	for candies != 0 {
		if candies - i < 0 {
			result[j] = result[j]+ candies
			candies = 0
		}else {
			result[j] = result[j] + i
			candies = candies - i
		}

		i = i + 1

		j = j + 1
		if j == numPeople {
			j = 0
		}
	}
	return result
}