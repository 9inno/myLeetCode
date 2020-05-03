package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Print(maxDiff(9))
}
func maxDiff(num int) int {
	str := strconv.Itoa(num)
	a := 0
	b := 0
	i := 0
	for i < len(str) {
		if str[i] == '9' {
			i++
		} else {
			a , _ = strconv.Atoi(strings.Replace(str,string(str[i]), "9",-1))
			break
		}
	}
	if a == 0{
		a = num
	}
	if str[0] != '1' {
		b , _ =strconv.Atoi(strings.Replace(str,string(str[0]), "1",-1))
	} else {
		i = 1
		for i < len(str) {
			if str[i] != '0' && str[i] != '1' {
				b, _ = strconv.Atoi(strings.Replace(str,string(str[i]), "0",-1))
				break
			}else  {
				i ++
				continue
			}
		}
		if b == 0 {
			b = num
		}
	}
	return a - b
}
