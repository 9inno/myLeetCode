package leetcode

import "strings"

func CountCharacters(words []string, chars string) int {

	charMap := make(map[string]int)
	result := 0
	for i := 0; i < len(chars); i++ {
		if _ , ok := charMap[string(chars[i])]; ok {
			charMap[string(chars[i])] ++
		} else {
			charMap[string(chars[i])] = 1
		}
	}

	for _, word := range words {
		wordMap := make(map[string]int)
		for j := 0; j < len(word); j ++ {
			if _ , ok := wordMap[string(word[j])]; ok {
				wordMap[string(word[j])] ++
			} else {
				wordMap[string(word[j])] = 1
			}
		}
		flag := true
		for chr , count := range wordMap{
			if _ , ok := charMap[chr]; ok{
				if charMap[chr] < count {
					flag = false
				}
			} else {
				flag = false
			}
		}
		if flag {
			result += len(word)
		}
	}
	return result

}


func CountCharacters1(words []string, chars string) int {
	result := 0
	loop : for _ , word := range words {
		for _ , letter := range word {
			if strings.Count(word, string(letter)) > strings.Count(chars, string(letter)) {
				continue loop
			}
		}
		result += len(word)
	}
	return result
}

