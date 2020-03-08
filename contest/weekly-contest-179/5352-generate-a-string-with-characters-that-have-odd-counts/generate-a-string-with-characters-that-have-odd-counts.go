package leetcode

func generateTheString(n int) string {
	result := ""
	if n % 2 == 0 {
		for i := 0; i < n - 1 ; i = i + 1  {
			result = result + "a"
		}
		result = result + "b"
	}else {
		for i := 0; i < n  ; i = i + 1  {
			result = result + "a"
		}
	}
	return result
}