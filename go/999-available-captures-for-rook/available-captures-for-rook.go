package leetcode



func NumRookCaptures(board [][]byte) int {

	result := 0
	outLoop : for i , row := range board {
		for j , column := range row {

			if string(column) == "R" {
				// 上
				for t := i; t > 0; t-- {
					if string(board[t][j]) == "B" {
						break
					}
					if string(board[t][j]) == "p" {
						result ++
						break
					}
				}
				// 下
				for b := i; b < len(board) ; b ++ {
					if string(board[b][j]) == "B" {
						break
					}
					if string(board[b][j]) == "p" {
						result ++
						break
					}
				}
				// 左
				for l := j ; l > 0 ; l -- {
					if string(board[i][l]) == "B" {
						break
					}
					if string(board[i][l]) == "p" {
						result ++
						break
					}
				}
				// 右
				for r := j; r < len(row); r++ {
					if string(board[i][r]) == "B" {
						break
					}
					if string(board[i][r]) == "p" {
						result ++
						break
					}
				}
				break outLoop
			}
		}
	}

	return result
}