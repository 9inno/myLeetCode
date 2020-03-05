package leetcode


func ReverseString(s []byte)  {
	i , j := 0 , len(s) - 1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i = i + 1
		j= j - 1
	}
}

func ReverseString1(s []byte)  {
	type r func(int , int)
	var recursion r
	recursion = func (left int , right int) {
		if left < right {
			s[left], s[right] = s[right], s[left]
			recursion(left + 1, right -1)
		}

	}
	recursion(0, len(s) - 1)
}

