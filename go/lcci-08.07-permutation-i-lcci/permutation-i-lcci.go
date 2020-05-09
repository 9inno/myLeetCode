package main

func main() {
	permutation("asd")
}


func permutation(S string) []string {
	result := []string{}
	recursion([]byte(S), &result, 0, len(S))
	return result
}

func recursion(arr []byte, result *[]string, n int, count int) {
	if count == 0 {
		return
	}
	if count == n {
		*result = append( *result, string(arr))
	}
	for i := n ; i < count; i ++ {
		arr[n], arr[i] = arr[i] , arr[n]
		recursion(arr, result, n + 1, count)
		arr[n], arr[i] = arr[i] , arr[n]
	}
}