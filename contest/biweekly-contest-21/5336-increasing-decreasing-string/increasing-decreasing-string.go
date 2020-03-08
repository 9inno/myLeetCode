package main

import (
	"fmt"
	"sort"
)

func main() {
	a:= sortString("aaaabbbbcccc")
	fmt.Print(a)
}

func sortString(s string) string {
	var sSlice []int
	var result string = ""
	for _ , value := range s {
		sSlice = append(sSlice , int(value))
	}

	for len(sSlice) != 0 {
		sort.Ints(sSlice)
		lastAppend := 0
		i := 0
		for i < len(sSlice){
			if sSlice[i] != lastAppend {
				result = result + string(sSlice[i])
				lastAppend = sSlice[i]
				if i < len(sSlice) - 1 {
					sSlice = append(sSlice[:i], sSlice[i + 1:]...)
				}else {
					sSlice = sSlice[:i]
				}
			}else {
				i = i + 1
			}
		}
		lastAppend = 0
		i = 0
		sort.Sort(sort.Reverse(sort.IntSlice(sSlice)))
		for i < len(sSlice){
			if sSlice[i] != lastAppend {
				result = result + string(sSlice[i])
				lastAppend = sSlice[i]
				if i < len(sSlice) - 1 {
					sSlice = append(sSlice[:i], sSlice[i + 1:]...)
				}else {
					sSlice = sSlice[:i]
				}
			}else {
				i = i + 1
			}
		}
	}
	return result

}