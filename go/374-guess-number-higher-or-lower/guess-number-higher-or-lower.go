package leetcode


/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */
var Target int
func guess(a int) int {
	if a > Target {
		return -1
	}
	if a <Target {
		return 1
	}
	return 0
}


func GuessNumber(n int) int {
	start := 1
	end := n

	for start < end {
		mid := start + (end - start) /2
		if guess(mid) == 0 {
			return mid
		} else  if guess(mid) == 1 {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}
