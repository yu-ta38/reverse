package main

import (
	"fmt"
	"time"
)

const GRID = 8
const LIMIT = 600.0

//const LIMIT = 10.0

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
	//b.tokens[3][4] = 1
	//b.tokens[4][3] = 1
	b.tokens[4][4] = 1

}

// 盤面を表示
func (b *Board) display() {
	fmt.Print(" ")
	for i := 0; i < GRID; i++ {
		fmt.Printf(" %d", i)
	}
	fmt.Print("\n")

	for i := 0; i < GRID; i++ {
		fmt.Print(i)
		for j := 0; j < GRID; j++ {
			if b.tokens[i][j] == 0 {
				fmt.Print(" .")
			} else if b.tokens[i][j] == 1 {
				fmt.Print(" O")
			} else if b.tokens[i][j] == -1 {
				fmt.Print(" X")
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

// 1つでも置ける場所があるかの判定
func (b *Board) check(u string) int {
	fmt.Print("ここに置けます: ")
	res := -1
	for i := 0; i < GRID; i++ {
		for j := 0; j < GRID; j++ {
			if b.tokens[i][j] != 0 {
				continue
			}
			r := b.reverse(i, j, u, false)
			if r > 0 {
				// (i, j)は置ける
				//return 1
				fmt.Printf("(%d, %d)", j, i)

				res = 1
			}
		}
	}

	fmt.Print("\n")
	return res
}

// 白と黒の駒の数を計算
func (b *Board) count() {
	O_count := 0
	X_count := 0

	for i := 0; i < GRID; i++ {
		for j := 0; j < GRID; j++ {
			if b.tokens[i][j] == 1 {
				O_count++
			} else if b.tokens[i][j] == -1 {
				X_count++
			}
		}
	}

	if X_count < O_count {
		fmt.Print("Oの勝利！\n")
	} else if X_count > O_count {
		fmt.Print("Xの勝利！\n")
	} else if X_count == O_count {
		fmt.Print("ひきわけ\n")
	}

	fmt.Printf("(O:%d / X:%d)\n", O_count, X_count)
}

func main() {
	var board Board
	var argv [2]int
	u := "O"
	player := 1
	_check := 0
	skip := 0
	var Otime time.Time
	var Xtime time.Time

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
		start := time.Now()
		// fmt.Scan(&argv[1], &argv[0])
		fmt.Scan(&argv[1])
		fmt.Scan(&argv[0])
		end := time.Now()
		elapsed := end.Sub(start)

		// fmt.Println(elapsed)
		//fmt.Printf("%T", elapsed)

		//fmt.Printf("%d m %.6f s\n", Xtime.Minute(), Xtime.Second())
		if player == 1 {
			// "O"の時間 ++
			Otime = Otime.Add(elapsed)
		} else if player == -1 {
			// "X"の時間 ++
			Xtime = Xtime.Add(elapsed)
		}
		fmt.Println("O ", Otime.Format("04:05"))
		fmt.Println("X ", Xtime.Format("04:05"))

		// 持ち時間チェック
		if Otime.Second() > LIMIT {
			fmt.Println("O time over")
			fmt.Print("試合終了〜\nX の勝ち！\n")
			return
		}
		if Xtime.Second() > LIMIT {
			fmt.Println("X time over")
			fmt.Print("試合終了〜\nO の勝ち！\n")
			return
		}

		// 盤面の外に置こうとした時にエラーを返す
		if argv[0] < 0 || argv[0] >= GRID || argv[1] < 0 || argv[1] >= GRID {
			fmt.Print("そこには置けません！\n\n")
			continue
		}

		// 既に駒が置かれていた時にエラーを返す
		if board.tokens[argv[0]][argv[1]] != 0 {
			fmt.Print("そこには置けません！\n\n")
			continue
		}

		r := board.reverse(argv[0], argv[1], u, true)
		if r == 0 {
			fmt.Print("そこには置けません！\n\n")
			continue
		}
		board.put(argv[0], argv[1], u)
		player = player * -1

		if player == 1 {
			u = "O"
		} else if player == -1 {
			u = "X"
		}
		fmt.Print("\n")
	}
	board.display()
	fmt.Print("試合終了〜\n")
	board.count()
}
