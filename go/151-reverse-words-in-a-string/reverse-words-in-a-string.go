package leetcode

import "strings"

func ReverseWords(s string) string {
	var slice []string
	word := ""
	for i := len(s) - 1 ; i >= 0 ; i -- {
		tmpStr := string(s[i])
		if tmpStr  == " " {
			if word != "" {
				slice = append(slice, word)
			}
			word = ""
			continue
		}
		word = tmpStr + word
	}
	if word != "" {
		slice = append(slice, word)
	}

	return strings.Join(slice, " ")
}
