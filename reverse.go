package main

import "fmt"

const GRID = 8

type Board struct {
	tokens [GRID][GRID]int
}

func (b *Board) init() {
	// 0で初期化
	for i := 0; i < GRID; i++ {
		for j := 0; j < GRID; j++ {
			b.tokens[i][j] = 0
		}
	}
}

func (b *Board) display() {
	for i := 0; i < GRID; i++ {
		for j := 0; j < GRID; j++ {
			if b.tokens[i][j] == 0 {
				fmt.Print(".")
			} else if b.tokens[i][j] == 1 {
				fmt.Print("O")
			} else if b.tokens[i][j] == 2 {
				fmt.Print("X")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	var board Board
	board.init()
	board.display()
}
