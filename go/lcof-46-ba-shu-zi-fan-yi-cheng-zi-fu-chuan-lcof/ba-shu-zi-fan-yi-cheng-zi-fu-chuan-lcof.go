package main

func main() {
	translateNum(25)
}

func translateNum(num int) int {
	result := 1
	pre := 0
	prepre := 0
	tmpNum := num % 10
	for num != 0 {
		tmp := num % 10
		prepre , pre , result = pre, result, 0
		result += pre
		if tmp * 10 + tmpNum <= 25 && tmp * 10 +tmpNum >= 10 {
			result += prepre
		}
		tmpNum = tmp
		num /= 10
	}

	return result
}
