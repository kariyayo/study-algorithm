package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const END = 8

type Queen struct {
	row       int
	col       int
	nextChild int
}

type Stack struct {
	xs  [10]Queen // xs[0]は使わない
	pos int
}

func (s *Stack) push(q Queen) {
	s.pos++
	s.xs[s.pos] = q
}

func (s *Stack) pop() Queen {
	q := s.xs[s.pos]
	s.pos--
	return q
}

func (s *Stack) size() int {
	return s.pos
}

func (s *Stack) ng(r, c int) bool {
	q := s.xs[r+1]
	if q.col == c {
		return true
	}
	return false
}

func (s *Stack) rowCheck(nextRow int) bool {
	// 次に置こうと思ってる行にはまだ置かれていないので必ずtrue
	return true
}

func (s *Stack) colCheck(nextCol int) bool {
	// 次に置こうと思ってる列に既にクイーンがいないかどうかをチェック
	for r := 0; r < s.pos; r++ {
		if s.ng(r, nextCol) {
			return false
		}
	}
	return true
}

func (s *Stack) diagonalCheck(row, col int) bool {
	// 次に置こうと思ってる場所の左上方向をチェック
	dr := row
	dc := col
	for dr > 0 && dc > 0 {
		dr--
		dc--
	}
	for dr < row && dc < col {
		if s.ng(dr, dc) {
			return false
		}
		dr++
		dc++
	}
	// 次に置こうと思ってる場所の右上方向をチェック
	dr = row
	dc = col
	for dr > 0 && dc < 7 {
		dr--
		dc++
	}
	for dr < row && dc > col {
		if s.ng(dr, dc) {
			return false
		}
		dr++
		dc--
	}
	// 問題ないのでOK
	return true
}

func solver(initial [][]int) {
	stack := &Stack{}
	var initRoot = fromInitial(initial, 0)
	for rootCol := 0; rootCol < END; rootCol++ {
		var root Queen
		if initRoot != nil {
			// 初期値で1行目が与えられている場合
			root = *initRoot
			rootCol = END
		} else {
			root = Queen{row: 0, col: rootCol}
		}

		stack.push(root)
		for stack.size() > 0 {
			q := stack.pop()
			if q.nextChild == END {
				// 次に挑戦する場所がないのでさらにpop
				continue
			}

			// 次にチャレンジする行のクイーン
			row := q.row + 1
			var qq *Queen
			qq = fromInitial(initial, row)
			if qq != nil {
				q.nextChild = END // 初期値から取得した場合は他の列にチャレンジしないので固定値をセット
			} else {
				qq = &Queen{row: row, col: q.nextChild}
				q.nextChild++
			}
			next := *qq

			stack.push(q)

			// チャレンジ場所にクイーンを置けるかチェック
			if !stack.rowCheck(next.row) ||
				!stack.colCheck(next.col) ||
				!stack.diagonalCheck(next.row, next.col) {
				continue
			}

			stack.push(next)

			// 全部置くことができたので完了
			if stack.size() == 8 {
				rootCol = END
				break
			}
		}
	}

	printAnswer(stack)
}

func fromInitial(initial [][]int, row int) *Queen {
	for col, v := range initial[row] {
		if v == 1 {
			return &Queen{row: row, col: col}
		}
	}
	return nil
}

func printAnswer(stack *Stack) {
	ss := make([]string, 8)
	for i := 7; stack.size() > 0; i-- {
		q := stack.pop()
		s := [8]byte{'.', '.', '.', '.', '.', '.', '.', '.'}
		s[q.col] = 'Q'
		ss[i] = fmt.Sprintf("%s", s)
	}
	for _, s := range ss {
		fmt.Println(s)
	}
}

func initCells(cells [][]int, initial [][]int) {
	for r, _ := range cells {
		for c, _ := range cells[r] {
			cells[r][c] = initial[r][c]
		}
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	initial := make([][]int, 8)
	for r, _ := range initial {
		initial[r] = make([]int, 8)
	}
	n := nextInt(sc)
	for i := 0; i < n; i++ {
		r := nextInt(sc)
		c := nextInt(sc)
		initial[r][c] = 1
	}
	solver(initial)
}

// ----------

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}

func nextNumber(sc *bufio.Scanner) float64 {
	sc.Scan()
	f, err := strconv.ParseFloat(sc.Text(), 32)
	if err != nil {
		panic(err)
	}
	return f
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func printArray(xs []int) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
