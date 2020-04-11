package leetcode


func SuperEggDrop(K int, N int) int {
	if K == 1 {
		return N
	}
	if N == 0 {
		return 0
	}
	var (
		dp = make([][]int,K+1)
		m = 0
	)
	for i := 0; i < K+1; i ++ {
		dp[i] = make([]int,N+1)
	}
	for dp[K][m] != N {
		m ++
		for k := 1; k < K+1; k ++ {
			dp[k][m] = dp[k][m-1] + dp[k-1][m-1] + 1
			if dp[k][m] > N {
				return m
			}
		}
	}
	return m
}


func SuperEggDrop1(K int, N int) int {
	N++
	for base := 0; ; base++ {
		maxN := 0
		for i := 0; i <= K && i <= base; i++ {
			maxN += comb(i, base)
		}
		if maxN >= N {
			return base
		}
	}
}

func comb(m, n int) int {
	if m > n {
		m = n
	}
	if m > n/2 {
		m = n - m
	}
	ans := 1
	for i := n - m + 1; i <= n; i++ {
		ans *= i
	}
	for i := 2; i <= m; i++ {
		ans /= i
	}
	return ans
}

