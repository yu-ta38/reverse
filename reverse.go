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
	b.tokens[3][3] = 1
	b.tokens[3][4] = 2
	b.tokens[4][3] = 2
	b.tokens[4][4] = 1
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

func (b *Board) put(x, y int, u string) {
	if u == "O" {
		b.tokens[x][y] = 1
	} else if u == "X" {
		b.tokens[x][y] = 2
	}
}

//func (b *Board) get(x, y int) string {
//	get_num := ""
//	if b.tokens[x][y] == 0 {
//		get_num = "."
//	} else if b.tokens[x][y] == 1 {
//		get_num = "O"
//	} else if b.tokens[x][y] == 2 {
//		get_num = "X"
//	}
//	return get_num
//}

func (b *Board) reverse(x, y int, u string) {
	// 探索の方向
	dx := [](int){1, 0, -1, 0, 1, 1, -1, -1}
	dy := [](int){0, 1, 0, -1, 1, -1, 1, -1}
	var player, enemy int

	if u == "O" {
		player = 1
		enemy = 2
	} else if u == "X" {
		player = 2
		enemy = 1
	}

	for i := 0; i < GRID; i++ {
		ix := x + dx[i]
		iy := y + dy[i]
		for {
			if ix < 0 || ix >= GRID || iy < 0 || iy >= GRID {
				break
			}
			if b.tokens[ix][iy] == 0 {
				break
			} else if b.tokens[ix][iy] == player {
				break
			} else if b.tokens[ix][iy] == enemy {
				ix += dx[i]
				iy += dy[i]
			}
		}
		for !(ix == x && iy == y) {
			ix -= dx[i]
			iy -= dy[i]
			b.tokens[ix][iy] = player
			// 			ix -= dx[i]
			// 			iy -= dy[i]
		}
	}
}

func main() {
	var board Board
	board.init()

	board.put(3, 5, "O")
	board.reverse(3, 5, "O")
	//	board.put(1, 1, "X")

	board.display()
}
