package leetcode

import (
	"sort"
	"strings"
)

func MinimumLengthEncoding(words []string) int {
	if words==nil || len(words)==0{
		return 0
	}
	for i:=0;i<len(words);i++{
		reverse(&words[i])
	}
	sort.Strings(words)
	words=append(words,"*")
	res:=0
	for i:=0;i<len(words)-1;i++{
		if strings.HasPrefix(words[i+1],words[i]){
			continue
		}
		res+=len(words[i])+1
	}
	return res
}

func reverse(str *string) {
	s := []rune(*str)
	right := len(*str) - 1
	left := 0
	for right > left {
		s[left] , s[right] = s[right] , s[left]
		left ++
		right --
	}
	*str = string(s)
}
