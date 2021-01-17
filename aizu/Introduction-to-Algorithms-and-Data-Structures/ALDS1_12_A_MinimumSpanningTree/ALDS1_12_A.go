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

type Edge struct {
	w int
	s int
	t int
}

func (e *Edge) String() string {
	return fmt.Sprintf("{w:%v s:%v t:%v}", e.w, e.s, e.t)
}

func merge(n int, left, right []*Edge) []*Edge {
	i := 0
	j := 0
	k := 0
	result := make([]*Edge, n)
	for {
		if i >= len(left) && j >= len(right) {
			break
		}
		var next *Edge
		if i < len(left) && j >= len(right) {
			next = left[i]
			i++
		} else if i >= len(left) && j < len(right) {
			next = right[j]
			j++
		} else {
			if left[i].w < right[j].w {
				next = left[i]
				i++
			} else {
				next = right[j]
				j++
			}
		}
		result[k] = next
		k++
	}
	return result
}

func sort(es []*Edge) []*Edge {
	n := len(es)
	if n <= 1 {
		return es
	}
	left := sort(es[0 : n/2])
	right := sort(es[n/2 : n])
	return merge(n, left, right)
}

func solver(n int, es []*Edge) {
	sorted := sort(es)
	uf := NewUnionFind(n)
	sum := 0
	for _, edge := range sorted {
		if uf.Same(edge.s, edge.t) {
			continue
		}
		uf.Unite(edge.s, edge.t)
		sum += edge.w
	}
	fmt.Println(sum)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	es := make([]*Edge, n*n)
	k := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			w := nextInt(sc)
			if w != -1 {
				edge := &Edge{w: w, s: i, t: j}
				es[k] = edge
				k++
			}
		}
	}
	solver(k, es[0:k])
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

func printArray(xs []*Edge) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
