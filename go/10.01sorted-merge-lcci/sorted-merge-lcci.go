package leetcode


func Merge(A []int, m int, B []int, n int)  {
	p1 := m - 1
	p2 := n -1
	current := m + n - 1
	for p1 >= 0 && p2 >= 0 {
		if A[p1] < B[p2] {
			A[current] = B[p2]
			p2 = p2 - 1
		} else {
			A[current] = A[p1]
			p1 = p1 - 1
		}
		current = current - 1
	}
	if p2 != -1 {
		copy(A, B[0 : p2 + 1])
	}
}