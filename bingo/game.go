package bingo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	Numbers []int64
	Boards  []Board
	Wins    map[int]int64
	Order   []int
}

type Win struct {
	Board int
	Score int64
}

func New() Game {
	var Numbers []int64
	var Boards []Board
	var Order []int
	Wins := make(map[int]int64)

	return Game{Numbers, Boards, Wins, Order}
}

func (game *Game) Run() int64 {
	for _, number := range game.Numbers {
		for i := range game.Boards {
			_, ok := game.Wins[i]
			if ok {
				continue
			}

			bingoSum := game.Boards[i].Mark(number)

			if bingoSum != -1 {
				game.Wins[i] = game.Boards[i].unmarkedSum() * number
				game.Order = append(game.Order, i)
			}
		}
	}

	return -1
}

func FromInput(input string) Game {
	f, err := os.Open(fmt.Sprintf("internal/inputs/%s.txt", input))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	game := New()

	s := bufio.NewScanner(f)

	// Read the game numbers
	s.Scan()
	for _, s := range strings.Split(s.Text(), ",") {
		n, _ := strconv.Atoi(s)
		game.Numbers = append(game.Numbers, int64(n))
	}

	// New line here means that there's another board
	for s.Scan() {
		board := NewBoard()

		for i := 0; i < 5; i++ {
			s.Scan()

			j := 0

			for _, s := range strings.Split(s.Text(), " ") {
				n, err := strconv.Atoi(s)

				if err == nil {
					board.Numbers[i][j] = int64(n)
					j++
				}
			}

		}

		game.Boards = append(game.Boards, board)
	}

	return game
}
