package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UnionFind struct {
	parents []int
	size    []int
}

func (u *UnionFind) root(x int) int {
	if u.parents[x] == -1 {
		return x
	}
	p := u.root(u.parents[x])
	// 親を根で置き換えてしまう
	//   経路圧縮で効率化 O(N) -> O(logN)
	u.parents[x] = p
	return p
}

func (u *UnionFind) Unite(x, y int) {
	xRoot := u.root(x)
	yRoot := u.root(y)
	// 既に同じグループの場合は何もしない
	if xRoot == yRoot {
		return
	}
	// サイズが小さい方のグループを大きい方のグループに統合する
	//   union by size による効率化 O(logN) -> O(α(N))
	if u.size[xRoot] > u.size[yRoot] {
		u.parents[yRoot] = xRoot
		u.size[xRoot] = u.size[xRoot] + u.size[yRoot]
	} else {
		u.parents[xRoot] = yRoot
		u.size[yRoot] = u.size[yRoot] + u.size[xRoot]
	}
}

func (u *UnionFind) Same(x, y int) bool {
	return u.root(x) == u.root(y)
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = -1
	}
	// sizeの大小比較ができれば良いので特に初期化していないが、
	// 値自体に意味がある場合は 1 で初期化する
	size := make([]int, n)
	return &UnionFind{parents: parents, size: size}
}

const uniteCommand = 0
const sameCommand = 1

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	uf := NewUnionFind(n)
	q := nextInt(sc)
	for i := 0; i < q; i++ {
		command := nextInt(sc)
		p1 := nextInt(sc)
		p2 := nextInt(sc)
		switch command {
		case uniteCommand:
			uf.Unite(p1, p2)
		case sameCommand:
			isSame := uf.Same(p1, p2)
			if isSame {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		default:
			panic("not supported command")
		}
	}
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
