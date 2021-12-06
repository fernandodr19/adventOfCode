package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// read file to buffer
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}

	// convert content to a string
	str := string(b)

	// parse input
	numbers, boards, err := parseInput(str)
	if err != nil {
		panic(err)
	}

	fmt.Println(getWinnerBoardMagicNumber(numbers, boards)) // 54275
	fmt.Println(getLoserBoardMagicNumber(numbers, boards))  // 13158
}

// https://adventofcode.com/2021/day/4#part1
func getWinnerBoardMagicNumber(numbers []int, boards []Board) int {
	for _, n := range numbers {
		for _, b := range boards {
			b.mark(n)
			if b.isCompleted() {
				fmt.Println("BINGO")
				return b.unmarkedSum() * n
			}
		}
	}
	return -1
}

// https://adventofcode.com/2021/day/4#part2
func getLoserBoardMagicNumber(numbers []int, boards []Board) int {
	completedBoardSet := make(map[int]bool)
	for _, n := range numbers {
		for j, b := range boards {
			_, ok := completedBoardSet[j]
			if ok {
				// already completed
				continue
			}
			b.mark(n)
			if b.isCompleted() {
				completedBoardSet[j] = true
				if len(completedBoardSet) == len(boards) {
					fmt.Println("LOSER")
					return b.unmarkedSum() * n
				}
			}
		}
	}
	fmt.Println(completedBoardSet)
	return -1
}

// parseInput parses the .txt input into numbers and boards
func parseInput(input string) ([]int, []Board, error) {
	numbers := make([]int, 0)
	boards := make([]Board, 0)
	var err error

	elements := strings.Split(input, "\n\n")
	for i, element := range elements {
		if i == 0 {
			// first element which are the numbers
			numbers, err = parseNumbers(element)
			if err != nil {
				return nil, nil, err
			}
			continue
		}

		// when execution flow reaches this point every element is representing a board
		board, err := parseBoard(element)
		if err != nil {
			return nil, nil, err
		}

		boards = append(boards, board)
	}

	return numbers, boards, nil
}

// parseNumbers extract the bingo numbers from the input
func parseNumbers(commaSeparatedNumbers string) ([]int, error) {
	numbers := make([]int, 0)
	for _, nStr := range strings.Split(commaSeparatedNumbers, ",") {
		n, err := strconv.Atoi(nStr)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}

	return numbers, nil
}

// parseBoard extracts the boards from the input
func parseBoard(rawBoard string) (Board, error) {
	boardSize := strings.Count(rawBoard, "\n") + 1

	var bingoBoard Board
	// init board y
	bingoBoard = make([][]BoardElement, boardSize)

	for i, rawBoardLine := range strings.Split(rawBoard, "\n") {
		// init board x
		bingoBoard[i] = make([]BoardElement, boardSize)

		// remove double spaces
		rawBoardLine = strings.ReplaceAll(rawBoardLine, "  ", " ")
		// remove trailing spaces
		rawBoardLine = strings.Trim(rawBoardLine, " ")
		// split them by space
		lineElement := strings.Split(rawBoardLine, " ")

		for j, lineElement := range lineElement {
			n, err := strconv.Atoi(lineElement)
			if err != nil {
				return nil, err
			}
			bingoBoard[i][j] = BoardElement{
				number: n,
				marked: false,
			}
		}
	}

	return bingoBoard, nil
}
