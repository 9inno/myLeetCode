package main

import (
	"fmt"
	"sort"
)

func main() {
	a:= "bxfowqvnrhuzwqohquamvszkvunb"
	b:= "xjegbjccjjxfnsiearbsgsofywtq"

	fmt.Print(checkIfCanBreak(a, b))
}

func checkIfCanBreak(s1 string, s2 string) bool {
	slice1 := []byte(s1)
	slice2 := []byte(s2)
	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})
	sort.Slice(slice2, func(i, j int) bool {
		return slice2[i] < slice2[j]
	})
	var flag bool
	j :=0

	for {
		if slice1[j] == slice2[j]{
			j++
			continue
		}
		flag = slice1[j] >= slice2[j]
		break
	}

	for i := j; i < len(s1); i ++ {
		if flag {
			if slice1[i] < slice2[i] {
				return false
			}
		} else  {
			if slice1[i] > slice2[i] {
				return false
			}
		}

	}
	return true
}
