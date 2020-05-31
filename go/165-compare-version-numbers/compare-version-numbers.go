package leetcode

import (
	"strconv"
	"strings"
)

func CompareVersion(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")
	len1 := len(arr1)
	len2 := len(arr2)
	if len1 > len2 {
		for i := 0 ; i < len1 - len2; i ++ {
			arr2 = append(arr2, "0")
		}
	} else {
		for i := 0 ; i < len2 - len1; i ++ {
			arr1 = append(arr1, "0")
		}
	}

	for i := 0; i < len(arr1); i ++ {
		tmp1, _ := strconv.Atoi(arr1[i])
		tmp2, _ := strconv.Atoi(arr2[i])
		if tmp1 > tmp2 {
			return 1
		} else if tmp1 < tmp2 {
			return -1
		}
	}
	return 0
}
