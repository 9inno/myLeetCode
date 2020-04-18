package main

import "sort"

func main() {
	getHappyString(1 ,3)
}


func getHappyString(n int, k int) string {
	result := []string{}
	hashMap := map[string]bool{}
	dp("a", n, "", &result, hashMap)
	dp("b", n, "", &result, hashMap)
	dp("c", n, "", &result, hashMap)
	if k > len(result) - 1 {
		return ""
	}
	sort.Strings(result)
	return result[k]

}

func dp(pre string, n int, tmpString string, result *[]string, hashMap map[string]bool) {

	if len(tmpString) == n {
		if _ , ok := hashMap[tmpString]; !ok {
			*result = append(*result , tmpString)
			hashMap[tmpString] = true
		}
		return
	}
	element := []string{"a", "b", "c"}

	for _ , e := range element {
		if e == pre {
			continue
		}
		dp(e, n, tmpString + e, result, hashMap)
	}
}
