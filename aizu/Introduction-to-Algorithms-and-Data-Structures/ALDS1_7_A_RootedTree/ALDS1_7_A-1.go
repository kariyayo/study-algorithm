package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	id       int
	parent   *Node
	children []*Node
}

func (n *Node) depth() int {
	d := 0
	p := n.parent
	for p != nil {
		d++
		p = p.parent
	}
	return d
}

func (n *Node) String() string {
	parentID := -1
	if n.parent != nil {
		parentID = n.parent.id
	}
	typeS := "internal node"
	if n.parent == nil {
		typeS = "root"
	} else if len(n.children) == 0 {
		typeS = "leaf"
	}
	ss := make([]string, len(n.children))
	for i, c := range n.children {
		ss[i] = fmt.Sprintf("%d", c.id)
	}
	return fmt.Sprintf("node %d: parent = %d, depth = %d, %s, [%v]", n.id, parentID, n.depth(), typeS, strings.Join(ss, ", "))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	ns := make([]*Node, n)
	for ni := 0; ni < n; ni++ {
		nodeID := nextInt(sc)
		k := nextInt(sc)
		children := make([]*Node, k)
		if ns[nodeID] == nil {
			ns[nodeID] = &Node{id: nodeID, parent: nil, children: children}
		}
		for ci := 0; ci < k; ci++ {
			childID := nextInt(sc)
			if ns[childID] == nil {
				ns[childID] = &Node{id: childID}
			}
			ns[childID].parent = ns[nodeID]
			children[ci] = ns[childID]
		}
		ns[nodeID].children = children
	}
	for _, node := range ns {
		fmt.Println(node)
	}
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
