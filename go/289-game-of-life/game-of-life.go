package main

func main() {
	a := [][]int {
		[]int{0,1,0},
		[]int{0,0,1},
		[]int{1,1,1},
		[]int{0,0,0},
	}

	gameOfLife(a)
}

func gameOfLife(board [][]int)  {
	var result [][]int

	for i , row := range board {
		var newRow []int
		for j , column := range row {
			liveCount := 0
			if i > 0 {
				if board[i - 1][j] == 1 {
					liveCount ++
				}
			}

			if i < len(board) - 1 {
				if board[i + 1][j] == 1 {
					liveCount ++
				}
			}

			if j > 0 {
				if board[i][j - 1] == 1 {
					liveCount ++
				}
			}
			if j < len(row) - 1 {
				if board[i][j + 1] == 1 {
					liveCount ++
				}
			}

			if i > 0 && j > 0 {
				if board[i - 1][j - 1] == 1 {
					liveCount ++
				}
			}

			if i < len(board) - 1  && j < len(row) - 1 {
				if board[i + 1][j + 1] == 1 {
					liveCount ++
				}
			}

			if j > 0 && i <len(board) - 1 {
				if board[i + 1][j - 1] == 1{
					liveCount ++
				}
			}

			if i > 0 && j < len(row) - 1 {
				if board[i - 1][j + 1] == 1 {
					liveCount ++
				}
			}

			if liveCount < 2 && column == 1 {
				newRow = append(newRow, 0)
				continue
			}
			if (liveCount == 2 || liveCount == 3) && column == 1 {
				newRow = append(newRow, 1)
				continue
			}
			if liveCount > 3 && column == 1 {
				newRow = append(newRow, 0)
				continue
			}

			if liveCount == 3 && column == 0 {
				newRow = append(newRow, 1)
				continue
			}

			newRow = append(newRow, column)
		}

		result = append(result, newRow)
	}
	copy(board,result)
}