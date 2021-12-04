package bingo

type Board struct {
	Numbers [5][5]int64
	Marked  [5][5]bool
}

func NewBoard() Board {
	return Board{
		Numbers: [5][5]int64{},
		Marked: [5][5]bool{
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
			{false, false, false, false, false},
		},
	}
}

func (board *Board) Mark(number int64) int64 {
	for i, row := range board.Numbers {
		for j, n := range row {
			if n == number {
				board.Marked[i][j] = true

				if board.isBingo(i, j) {
					return board.unmarkedSum()
				} else {
					return -1
				}
			}
		}
	}

	return -1
}

func (board *Board) isBingo(row, col int) bool {
	bingoRow, bingoCol := true, true

	for i := 0; i < 5; i++ {
		bingoCol = bingoCol && board.Marked[i][col]
	}

	for j := 0; j < 5; j++ {
		bingoRow = bingoRow && board.Marked[row][j]
	}

	return bingoRow || bingoCol
}

func (board *Board) unmarkedSum() int64 {
	var sum int64 = 0

	for i, row := range board.Marked {
		for j, marked := range row {
			if !marked {
				sum += board.Numbers[i][j]
			}
		}
	}

	return sum
}
