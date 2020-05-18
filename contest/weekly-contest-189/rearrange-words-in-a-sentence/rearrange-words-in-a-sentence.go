package main

import (
	"strings"
)

func main() {
	arrangeWords("  ")
}

func arrangeWords(text string) string {
	tmp := strings.Fields(strings.TrimSpace(strings.ToLower(text)))
	if len(tmp) == 0 {
		return ""
	}
	tmp = mergeSort(tmp, 0, len(tmp) - 1)
	result := strings.Join(tmp, " ")
	for _ , value := range result {
		result = strings.ToUpper(string(value)) + result[1:]
		break
	}
	return result
}

func mergeSort(text []string, start, end int) []string {
	if start >= end {
		return []string{text[start]}
	}
	mid := start + (end - start) /2
	left := mergeSort(text, start, mid)
	right := mergeSort(text, mid + 1 , end)
	var tmp []string
	i , j := 0, 0

	for i < len(left) && j < len(right) {
		if len(left[i]) <= len(right[j]) {
			tmp = append(tmp, left[i])
			i ++
		} else {
			tmp = append(tmp, right[j])
			j ++
		}
	}
	for i < len(left){
		tmp = append(tmp, left[i])
		i ++
	}
	for j < len(right) {
		tmp = append(tmp, right[j])
		j ++
	}
	return tmp
}