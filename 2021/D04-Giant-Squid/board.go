package main

type BoardElement struct {
	number int
	marked bool
}

type Board [][]BoardElement

func (b Board) mark(number int) {
	for i := range b {
		for j := range b[i] {
			if b[i][j].number == number {
				b[i][j].marked = true
			}
		}
	}
}

func (b Board) isCompleted() bool {
	for i := range b {
		fullRow, fullCol := true, true
		for j := range b[i] {
			if !b[i][j].marked {
				fullRow = false
			}

			if !b[j][i].marked {
				fullCol = false
			}
		}
		if fullRow || fullCol {
			return true
		}
	}

	return false
}

func (b Board) unmarkedSum() int {
	sum := 0
	for i := range b {
		for j := range b[i] {
			if !b[i][j].marked {
				sum += b[i][j].number
			}
		}
	}
	return sum
}
