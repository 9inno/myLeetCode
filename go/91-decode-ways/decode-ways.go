package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numDecodings("226"))
}
func numDecodings(s string) int {
	result := 0
	dp(s, &result)
	return result
}

func dp(s string, result *int) {
	if len(s) == 0 {
		*result ++
		return
	}
	if len(s) >= 2 {
		if value, err :=strconv.Atoi(s[0:2]); err == nil   {
			if value <= 26 {
				dp(s[2:], result)
			}
		}
		dp(s[1:], result)
	}else {
		dp(s[1:], result)
	}

}