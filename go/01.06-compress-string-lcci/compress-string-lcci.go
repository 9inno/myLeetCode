package leetcode

import "strconv"

func CompressString(S string) string {
	lenS := len(S)
	tmpCount := 0
	result := ""
	tmpString := ""
	for i := 0; i < lenS; i++ {

		if tmpString == "" {
			tmpString = string(S[i])
			result += string(S[i])
			tmpCount ++
		} else  {
			if ok := tmpString != string(S[i]); !ok {
				tmpCount ++
			} else {
				result += strconv.Itoa(tmpCount)
				result += string(S[i])
				tmpString = string(S[i])
				tmpCount = 1
			}
		}
	}
	result += strconv.Itoa(tmpCount)
	if len(result) >= lenS {
		return S
	}

	return result
}