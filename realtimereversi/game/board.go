package game

import (
	"fmt"
)

type Board struct {
    Cells [][]Color
}

// 盤面作成
func NewBoard() *Board {
    b := &Board{
        Cells: make([][]Color, 10)
    }
    for i := 0; i < 10; i++ {
        b.Cells[i] = make([]Color, 10)
    }

    // 上辺
    for i := 0; i < 10; i++ {
        b.Cells[0][i] = Wall
    }

    // 左辺と右辺？
    for i := 0; i < 9; i++ {
        b.Cells[i][0] = Wall
        b.Cells[i][9] = Wall
    }

    // 下辺？
    for i := 0; i < 9; i++ {
        b.Cells[9][i] = Wall
    }

    b.Cells[4][4] = White
    b.Cells[5][5] = White
    b.Cells[5][4] = Black
    b.Cells[4][5] = Black

    return b
}

func (b *Board) PutStone(x int32, y int32, c Color) error {
    // セルに石を置けるかチェック
    if !b.CanPutStone(x, y, c) {
        return fmt.Errorf("Can not put stone x=%v, y=%v, color=%v", x, y, ColorToStr(c))
    }

    // セルに石を置く
    b.Cells[x][y] = c

    for dx := -1; dx <= 1; dx++ {
        for dy := -1; dy <= 1; dy++ {
            if dx == 0 && dy == 0 {
                continue
            }
            b.TurnStoneByDirection(x, y, c, int32(dx), int32(dy))
        }
    }

    return nil
}

// セルに石が置けるか判定
func CanPutStone(x int32, y int32, c Color) bool {
    if b.Cells[x][y] != Empty {
        return false
    }

    for dx := -1; dx <= 1; dx++ {
        for dy := -1; dy <= 1; dy++ {
            if dx == 0 && dy == 0 {
                continue
            }

            if b.CountTurnableStonesByDirection(x, y, c, int32(dx), int32(dy)) > 0 {
                return true
            }
        }
    }

    return false
}

func (b *Board) CountTurnableStonesByDirection(x int32, y int32,c Color, dx int32, dy int32) int {
    cnt := 0

    nx := x + dx
    ny := y + dy
    for {
        nc := b.Cells[nx][ny]

        // 壁か自分の石であればループ終了
        if nc != OpponentColor(c) {
            break
        }

        // 相手の石なので数え上げ
        cnt++

        nx += dx
        ny += dy
    }

    // その方向にある相手の石が0より大きく、その先に自分の石がある場合は数を返す
    if cnt > 0 && b.Cells[nx][ny] == c {
        return cnt
    }

    return 0
}

// ある方向の石をひっくり返す。ひっくり返しても良い場合だけ呼ぶ
func (b *Board) TurnStonesByDirection(x int32, y int32, c Color, dx int32, dy int32) {
    nx := x + dx
    ny := y + dy

    for {
        nc := b.Cells[nx][ny]

        if nc != OpponentColor(c) {
            break
        }

        b.Cells[nx][ny] == c

        nx += dx
        ny += dy
    }
}

// 盤面内である石の置けるセルの数を数える
func (b *Board) AvailableCellCount(c Color) int {
    cnt := 0

    for i := 1; i < 9; i++ {
        for j := 1; j < 9; j++ {
            if b.CanPutStone(int32(i), int32(j), c) {
                cnt++
            }
        }
    }
    return cnt
}

// 盤面内に置かれている石の数を数える
func (b *Board) Score(c Color) int {
    cnt := 0

    for i := 1; i < 9; i++ {
        for j := 1; j < 9; j++ {
            if b.Cells[i][j] != c {
                continue
            }
            cnt++
        }
    }
    return cnt
}

// 盤面内で石が置かれていないセルの数を数える
func (b *Board) Rest() int {
    cnt := 0

    for i := 1; i < 9; i++ {
        for j := 1; j < 9; j++ {
            if b.Cells[i][j] == Empty {
                cnt++
            }
        }
    }
    return cnt
}
