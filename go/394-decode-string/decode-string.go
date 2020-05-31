package leetcode

import "strconv"

func DecodeString(s string) string {
	hashMap := map[string]bool{
		"0" : true,
		"1" : true,
		"2" : true,
		"3" : true,
		"4" : true,
		"5" : true,
		"6" : true,
		"7" : true,
		"8" : true,
		"9" : true,
	}
	var stringStack []string
	var numberStack []string
	tmpNumber := ""
outLoop:for _ , value := range s {
	if _ , ok := hashMap[string(value)] ; ok {
		tmpNumber += string(value)
		continue
	}
	if value == '[' {
		numberStack = append(numberStack, tmpNumber)
		tmpNumber = ""
	}
	if value == ']' {
		tmpStr := ""
		for {
			popValue := pop(&stringStack)
			if popValue == "[" {
				stringStack = append(stringStack, generate(pop(&numberStack), tmpStr))
				continue outLoop
			}
			tmpStr = popValue + tmpStr
		}
	}
	stringStack = append(stringStack, string(value))
}
	result := ""
	for _ , value := range stringStack {
		result +=value
	}
	return result
}

func pop(str *[]string) string {
	if len(*str) == 0 {
		return  ""
	}
	tmp := (*str)[len(*str) - 1]
	*str = (*str)[ :len(*str) - 1]
	return tmp
}

func generate(num string, subStr string) string {
	n, _ := strconv.Atoi(num)
	tmp := ""
	for i:=0 ; i < n ; i ++ {
		tmp += subStr
	}
	return tmp
}