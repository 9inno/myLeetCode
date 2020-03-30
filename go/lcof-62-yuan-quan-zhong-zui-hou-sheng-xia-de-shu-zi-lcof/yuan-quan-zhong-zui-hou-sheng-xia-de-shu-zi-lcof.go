package leetcode

func LastRemaining(n int, m int) int {
	tmp := make([]int,n)
	for i:=0 ; i < n; i++ {
		tmp[i] = i
	}
	z := 1
	for len(tmp) != 1 {
		for j := 0; j < len(tmp); j++ {
			if len(tmp) == 1 {
				break
			}
			if z == m {
				tmp = append(tmp[:j], tmp[j+1:]...)
				z = 1
				j --
			}else {
				z ++
			}
		}
	}

	return tmp[0]
}

func LastRemaining1(n, m int) int {

	result := 0
	for i := 2 ;i < n + 1 ; i ++{
		result = (m + result) % i
	}
	return result
}