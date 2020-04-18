package main

func main() {
	findMinFibonacciNumbers(5)
}
func findMinFibonacciNumbers(k int) int {
	var closure  func(n int)
	fibo := []int{1, 1}
	result := 0
	for fibo[len(fibo)-1] < k {
		fibo = append(fibo,fibo[len(fibo)-1] + fibo[len(fibo)-2])
	}
	if k == fibo[len(fibo)-1] {
		return 1
	}
	closure = func(n int) {
		result ++
		for _, value := range fibo {
			if value == n {
				return
			}
		}
		closure(min(0 , len(fibo) + 1, n, fibo))
	}
	closure(k)
	return result
}

func min(start, end , n int, nums []int) int {
	if end - start == 1 {
		return n - nums[start]
	}
	if n > nums[(start+end)/2] {
		return min((start + end) / 2 , end, n, nums)
	}else {
		return min(start , (start + end) / 2 , n, nums)
	}
}

