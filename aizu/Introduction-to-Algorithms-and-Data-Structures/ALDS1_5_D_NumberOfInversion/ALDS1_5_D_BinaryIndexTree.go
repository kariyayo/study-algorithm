package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type binaryIndexedTree struct {
	nodes map[int]int
	n     int
}

func newBinaryIndexedTree(n int) *binaryIndexedTree {
	t := &binaryIndexedTree{
		n:     n,
		nodes: make(map[int]int),
	}
	return t
}

func (t *binaryIndexedTree) add(i int, x int) {
	for k := i; k < t.n+1; k = k + (k & -k) {
		t.nodes[k] += x
	}
}

func (t *binaryIndexedTree) sum(a, b int) int {
	if a > b {
		return 0
	}
	sumb := t._sum(b)
	suma := t._sum(a - 1)
	// fmt.Printf("sum(%v):%v\n", b, sumb)
	// fmt.Printf("sum(%v):%v\n", a-1, suma)
	return sumb - suma
}

func (t *binaryIndexedTree) _sum(m int) int {
	result := 0
	for k := m; k > 0; k = k - (k & -k) {
		result += t.nodes[k]
	}
	return result
}

func (t *binaryIndexedTree) String() string {
	ss := []string{}
	for k, v := range t.nodes {
		ss = append(ss, fmt.Sprintf("%v:%v", k, v))
	}
	return strings.Join(ss, " ")
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

	max := 0
	ans := 0
	tree := newBinaryIndexedTree(1000000000 + 1)
	for i := 0; i < n; i++ {
		x := nextInt(sc) + 1 // 0を受け入れるためインデックス値として+1しておく
		if x > max {
			max = x
		}
		tree.add(x, 1)
		// fmt.Printf("x:%v\n", x)
		// fmt.Printf("tree:%v\n", tree)
		ans += tree.sum(x+1, max)
		// fmt.Printf("ans:%v\n\n", ans)
	}
	fmt.Println(ans)
}
