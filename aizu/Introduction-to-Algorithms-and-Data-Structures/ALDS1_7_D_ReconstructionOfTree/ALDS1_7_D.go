package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id     int
	parent *Node
	left   *Node
	right  *Node
}

var pos = 0
var preOrder []int

func reconstruct(inOrder []int, parent *Node) *Node {
	if len(inOrder) == 0 {
		return nil
	}

	id := preOrder[pos]
	node := &Node{id: id, parent: parent}
	pos++

	rootIdx := 0
	for i, v := range inOrder {
		if id == v {
			rootIdx = i
			break
		}
	}
	node.left = reconstruct(inOrder[:rootIdx], node)
	node.right = reconstruct(inOrder[rootIdx+1:], node)
	if parent == nil {
		fmt.Printf("%d\n", node.id)
	} else {
		fmt.Printf("%d ", node.id)
	}
	return node
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	preOrder = make([]int, n)
	inOrder := make([]int, n)
	for i := 0; i < n; i++ {
		preOrder[i] = nextInt(sc)
	}
	for i := 0; i < n; i++ {
		inOrder[i] = nextInt(sc)
	}
	reconstruct(inOrder, nil)
}

// ----------

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
	fmt.Fprintf(os.Stderr, format, a...)
}
