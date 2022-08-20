package main

import (
	"collections"
	"os"
	"panic"
	"readers"
	"strconv"
	"strings"
)

func Day04Part1() int {
	file, err := os.Open("assets/day04.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	numbersStrings := strings.Split(lines[0], ",")
	numbers := make([]int, len(numbersStrings))

	for i, s := range numbersStrings {
		numbers[i], err = strconv.Atoi(s)
		panic.Check(err)
	}

	var boards []BingoBoard

	for i := 2; i < len(lines); i += 6 {
		board := BingoBoard{}

		for row := 0; row < BingoBoardSize; row++ {
			values := collections.FilterEmptyFromString(strings.Split(lines[i+row], " "))

			for col, value := range values {
				intVal, err := strconv.Atoi(value)
				panic.Check(err)
				board.numbers = append(board.numbers, BingoNumber{
					value:    intVal,
					row:      row,
					col:      col,
					isMarked: false,
				})
			}
		}

		boards = append(boards, board)
	}

	for _, number := range numbers {
		for _, board := range boards {
			if board.TryMarkNumber(number) && board.IsWinning() {
				sum := 0
				for _, bingoNumber := range board.numbers {
					if !bingoNumber.isMarked {
						sum += bingoNumber.value
					}
				}
				return sum * number
			}
		}
	}

	return 0
}

func Day04Part2() int {
	file, err := os.Open("assets/day04.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	numbersStrings := strings.Split(lines[0], ",")
	numbers := make([]int, len(numbersStrings))

	for i, s := range numbersStrings {
		numbers[i], err = strconv.Atoi(s)
		panic.Check(err)
	}

	var boards []BingoBoard

	for i := 2; i < len(lines); i += 6 {
		board := BingoBoard{}

		for row := 0; row < BingoBoardSize; row++ {
			values := collections.FilterEmptyFromString(strings.Split(lines[i+row], " "))

			for col, value := range values {
				intVal, err := strconv.Atoi(value)
				panic.Check(err)
				board.numbers = append(board.numbers, BingoNumber{
					value:    intVal,
					row:      row,
					col:      col,
					isMarked: false,
				})
			}
		}

		boards = append(boards, board)
	}

	for round, number := range numbers {
		for i, board := range boards {

			if boards[i].score > 0 {
				continue
			}

			if board.TryMarkNumber(number) && board.IsWinning() {
				sum := 0
				for _, bingoNumber := range board.numbers {
					if !bingoNumber.isMarked {
						sum += bingoNumber.value
					}
				}
				boards[i].score = sum * number
				boards[i].round = round
			}
		}
	}

	maxRound := 0
	maxRoundScore := 0
	for _, board := range boards {
		if board.round > maxRound {
			maxRound = board.round
			maxRoundScore = board.score
		}
	}

	return maxRoundScore
}

const BingoBoardSize = 5

type BingoBoard struct {
	numbers []BingoNumber
	round   int
	score   int
}

type BingoNumber struct {
	value    int
	row      int
	col      int
	isMarked bool
}

func (board BingoBoard) TryMarkNumber(number int) bool {
	ret := false
	for i, bingoNumber := range board.numbers {
		if bingoNumber.value == number {
			board.numbers[i].isMarked = true
			ret = true
		}
	}

	return ret
}

func (board BingoBoard) BingoNumbersByRow(row int) []BingoNumber {
	var out []BingoNumber

	for _, number := range board.numbers {
		if number.row == row {
			out = append(out, number)
		}
	}

	return out
}

func (board BingoBoard) BingoNumbersByCol(col int) []BingoNumber {
	var out []BingoNumber

	for _, number := range board.numbers {
		if number.col == col {
			out = append(out, number)
		}
	}

	return out
}

func (board BingoBoard) IsWinning() bool {
	for row := 0; row < BingoBoardSize; row++ {
		numbersRow := board.BingoNumbersByRow(row)
		isWinning := true

		for i := 0; i < BingoBoardSize; i++ {
			if numbersRow[i].isMarked == false {
				isWinning = false
				break
			}
		}

		if isWinning {
			return true
		}
	}

	for col := 0; col < BingoBoardSize; col++ {
		numbersCol := board.BingoNumbersByCol(col)
		isWinning := true

		for i := 0; i < BingoBoardSize; i++ {
			if numbersCol[i].isMarked == false {
				isWinning = false
				break
			}
		}

		if isWinning {
			return true
		}
	}

	return false
}
