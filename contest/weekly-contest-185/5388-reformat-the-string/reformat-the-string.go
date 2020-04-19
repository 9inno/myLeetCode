package main

import (
	"fmt"
	"math"
)

func main() {
	a:= "a0b1c2"
	fmt.Print(reformat(a))
}
func reformat(s string) string {
	nums := []string{}
	strings := []string{}
	for _ , letter := range s {
		if letter >=48 && letter <= 57 {
			nums = append(nums, string(letter))
		}
		if letter >= 97 && letter <= 122 {
			strings = append(strings, string(letter))
		}
	}

	if math.Abs(float64(len(nums) - len(strings))) > 1 {
		return ""
	}
	result := ""
	flag := true
	if len(nums) < len(strings) {
		flag = false
	}
	for len(nums) != 0 || len(strings) != 0 {
		if flag {
			result += nums[0]
			nums = nums[1:]
			flag = false
		} else {
			result += strings[0]
			strings = strings[1:]
			flag = true
		}
	}
	return result

}
