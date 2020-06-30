package leetcode

func findPoisonedDuration(timeSeries []int, duration int) int {
	length := len(timeSeries)
	if length == 0 {
		return 0
	}
	result := duration
	tmp := timeSeries[0] + duration - 1

	for i := 1; i < length; i ++ {
		if tmp >= timeSeries[i] {
			result += timeSeries[i] + duration - 1 - tmp

		} else {
			result += duration
		}
		tmp = timeSeries[i] + duration - 1
	}
	return result
}
