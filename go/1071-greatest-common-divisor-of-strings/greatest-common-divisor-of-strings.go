package leetcode

func GcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	length := gcd(len(str1), len(str2))
	return str1[0:length]

}

func gcd(a int, b int) int  {
	if b == 0 {
		return a
	} else {
		return gcd(b, a % b)
	}
}
