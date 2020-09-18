package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type segmentTree struct {
	nodes []int
	n     int
}

func newSegmentTree(size int, xs []int) *segmentTree {
	tree := &segmentTree{}

	// 葉の数
	tree.n = size

	// 全ノード数
	tree.nodes = make([]int, 2*size-1)

	// 葉をセット
	start := len(tree.nodes) - size
	for i := 0; i < size; i++ {
		tree.nodes[start+i] = xs[i]
	}

	// 登っていきながら値をセット
	for k := start - 1; k >= 0; k-- {
		left := tree.nodes[2*k+1]
		right := tree.nodes[2*k+2]
		if left > right {
			tree.nodes[k] = right
		} else {
			tree.nodes[k] = left
		}
	}

	return tree
}

func (t *segmentTree) update(i int, x int) {
	start := len(t.nodes) - t.n + i
	t.nodes[start] = x

	// 親
	k := start
	for k > 0 {
		k = (k - 1) / 2
		left := t.nodes[2*k+1]
		right := t.nodes[2*k+2]
		if left > right {
			t.nodes[k] = right
		} else {
			t.nodes[k] = left
		}
	}
}

func (t *segmentTree) query(a int, b int) int {
	start := len(t.nodes) - t.n
	return t._query(start+a, start+b, 0, start, start+t.n-1)
}

func (t *segmentTree) _query(a int, b int, k int, left int, right int) int {
	if b <= left || right <= a {
		return math.MaxInt32
	} else if a <= left && right <= b {
		return t.nodes[k]
	} else {
		l := t._query(a, b, 2*k+1, left, (left+right)/2)
		r := t._query(a, b, 2*k+2, (left+right)/2, right)
		if l < r {
			return l
		} else {
			return r
		}
	}
}

func (t *segmentTree) String() string {
	return strings.Trim(fmt.Sprint(t.nodes), "[]")
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	s := sc.Text()
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

/**
 * $ go run segment_tree.go
 * > 8
 * > 5 3 7 9 1 4 6 2
 * 1 3 1 3 7 1 2 5 3 7 9 1 4 6 2
 * 2 3 2 3 7 4 2 5 3 7 9 10 4 6 2
 * 2
 */
func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = nextInt(sc)
	}

	tree := newSegmentTree(n, xs)
	fmt.Println(tree)
	tree.update(4, 10)
	fmt.Println(tree)
	fmt.Println(tree.query(0, 7))
}
