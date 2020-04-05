package main

func main() {
	print(countLargestGroup(13))
}

func countLargestGroup(n int) int {
	hashMap := map[int]int{}
	for i := 1 ; i < n + 1; i ++ {
		tmp := i
		sum := 0
		for tmp != 0 {
			sum += tmp % 10
			tmp /= 10
		}
		if _ , ok := hashMap[sum] ; ok{
			hashMap[sum] ++
		}else  {
			hashMap[sum] = 1
		}
	}

	max := 0
	result:=0
	resultMap := map[int]int{}
	for _ , value := range hashMap{
		if _ , ok := resultMap[value] ; ok{
			resultMap[value] ++
		}else {
			resultMap[value] = 1
		}
	}

	for key , count := range resultMap {
		if key > max {
			max = key
			result = count
		}
	}
	return result
}