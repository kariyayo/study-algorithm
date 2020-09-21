package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type binaryIndexedTree struct {
	nodes []int
	n     int
}

func newBinaryIndexedTree(n int) *binaryIndexedTree {
	tree := &binaryIndexedTree{}
	tree.n = n
	tree.nodes = make([]int, n+1) // インデックス1から始めるので全ノード数はn+1。最初の要素は無視
	return tree
}

// iは1開始
func (t *binaryIndexedTree) add(i int, x int) {
	for k := i; k-1 < t.n; k = k + (k & -k) {
		t.nodes[k] += x
	}
}

// iは1開始, bはt.n以下
func (t *binaryIndexedTree) getSum(a int, b int) int {
	return t.sum(b) - t.sum(a-1)
}

// iは1開始
func (t *binaryIndexedTree) sum(i int) int {
	result := 0
	for k := i; k > 0; k = k - (k & -k) {
		result += t.nodes[k]
	}
	return result
}

func (t *binaryIndexedTree) String() string {
	return strings.Trim(fmt.Sprint(t.nodes[1:]), "[]")
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
	q := nextInt(sc)
	tree := newBinaryIndexedTree(n)
	for ; q > 0; q-- {
		command := nextInt(sc)
		if command == 0 {
			i := nextInt(sc)
			x := nextInt(sc)
			tree.add(i, x)
		} else {
			a := nextInt(sc)
			b := nextInt(sc)
			result := tree.getSum(a, b)
			fmt.Println(result)
		}
	}
}
