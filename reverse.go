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
	b.tokens[3][4] = -1
	b.tokens[4][3] = -1
	b.tokens[4][4] = 1

}

func (b *Board) display() {
	fmt.Print(" ")
	for i := 0; i < GRID; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Print("\n")

	for i := 0; i < GRID; i++ {
		fmt.Print(i)
		for j := 0; j < GRID; j++ {
			if b.tokens[i][j] == 0 {
				fmt.Print(".")
			} else if b.tokens[i][j] == 1 {
				fmt.Print("O")
			} else if b.tokens[i][j] == -1 {
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
		b.tokens[x][y] = -1
	}
}

// func (b *Board) get(x, y int) string {
// 	get_num := ""
// 	if b.tokens[x][y] == 0 {
// 		get_num = "."
// 	} else if b.tokens[x][y] == 1 {
// 		get_num = "O"
// 	} else if b.tokens[x][y] == -1 {
// 		get_num = "X"
// 	}
// 	return get_num
// }

// flag: true -> 反転する
// flag: false -> チェックだけ
func (b *Board) reverse(x, y int, u string, flag bool) int {
	// 探索の方向
	dx := [](int){1, 0, -1, 0, 1, 1, -1, -1}
	dy := [](int){0, 1, 0, -1, 1, -1, 1, -1}
	var player, enemy int
	r := 0

	// O ->  1
	// X -> -1

	if u == "O" {
		player = 1
		enemy = -1
	} else if u == "X" {
		player = -1
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
				ix -= dx[i]
				iy -= dy[i]
				for !(ix == x && iy == y) {
					if flag {
						b.tokens[ix][iy] = player
					}
					r += 1
					ix -= dx[i]
					iy -= dy[i]
				}
				break
			} else if b.tokens[ix][iy] == enemy {
				ix += dx[i]
				iy += dy[i]
			}
		}
	}

	return r
}

func (b *Board) check(u string) int {
	for i := 0; i < GRID; i++ {
		for j := 0; j < GRID; j++ {
			if b.tokens[i][j] != 0 {
				continue
			}
			r := b.reverse(i, j, u, false)
			if r > 0 {
				return 1
			}
		}
	}

	return -1
}

func (b *Board) count() {
	white_count := 0
	black_count := 0

	for i := 0; i < GRID; i++ {
		for j := 0; j < GRID; j++ {
			if b.tokens[i][j] == 1 {
				white_count++
			} else if b.tokens[i][j] == -1 {
				black_count++
			}
		}
	}

	if black_count > white_count {
		fmt.Print("Oの勝利！\n")
	} else if black_count < white_count {
		fmt.Print("Xの勝利！\n")
	} else if black_count == white_count {
		fmt.Print("ひきわけ\n")
	}

	fmt.Printf("(O:%d / X:%d)\n", white_count, black_count)
}

func main() {
	var board Board
	var argv [2]int
	u := "O"
	player := 1
	_check := 0
	skip := 0

	board.init()

	for {
		// 置けるかどうか
		_check = board.check(u)
		if _check == -1 {
			player = player * -1
			if player == 1 {
				u = "O"
			} else if player == -1 {
				u = "X"
			}
			skip += 1
			if skip >= 2 {
				break
			}
			continue
		}
		skip = 0

		// 標準入力
		board.display()
		fmt.Printf("[%s] Input > ", u)
		fmt.Scan(&argv[0], &argv[1])
		//fmt.Println(argv)

		// TODO: range check
		if board.tokens[argv[0]][argv[1]] != 0 {
			fmt.Print("そこには置けません！\n")
			continue
		}

		r := board.reverse(argv[0], argv[1], u, true)
		if r == 0 {
			fmt.Print("そこには置けません！\n")
			continue
		}
		board.put(argv[0], argv[1], u)
		player = player * -1

		if player == 1 {
			u = "O"
		} else if player == -1 {
			u = "X"
		}
	}
	board.display()
	fmt.Print("試合終了〜\n")
	board.count()
}
