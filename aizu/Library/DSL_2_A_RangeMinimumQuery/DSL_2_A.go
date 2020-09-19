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

func newSegmentTree(n int) *segmentTree {
	tree := &segmentTree{}

	// 葉の数
	// セグメントツリーは完全二分木なので指定された数以上になる最小の2の冪
	size := 1
	for size < n {
		size = size * 2
	}
	tree.n = size

	// 全ノード数
	s := 2*tree.n - 1

	tree.nodes = make([]int, s)
	for i := 0; i < s; i++ {
		tree.nodes[i] = math.MaxInt32
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

func (t *segmentTree) find(a int, b int) int {
	return t._find(a, b, 0, 0, t.n)
}

func (t *segmentTree) _find(a int, b int, k int, left int, right int) int {
	if b <= left || right <= a {
		return math.MaxInt32
	} else if a <= left && right <= b {
		return t.nodes[k]
	} else {
		l := t._find(a, b, 2*k+1, left, (left+right)/2)
		r := t._find(a, b, 2*k+2, (left+right)/2, right)
		if l < r {
			return l
		}
		return r
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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	tree := newSegmentTree(n)

	q := nextInt(sc)
	for k := 0; k < q; k++ {
		command := nextInt(sc)
		if command == 0 {
			i := nextInt(sc)
			x := nextInt(sc)
			tree.update(i, x)
		} else {
			s := nextInt(sc)
			t := nextInt(sc)
			result := tree.find(s, t+1)
			fmt.Println(result)
		}
	}
}
