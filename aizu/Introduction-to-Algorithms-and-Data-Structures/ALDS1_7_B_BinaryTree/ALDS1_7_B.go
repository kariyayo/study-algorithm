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

func (n *Node) depth() int {
	d := 0
	p := n.parent
	for p != nil {
		d++
		p = p.parent
	}
	return d
}

func (n *Node) height() int {
	lHeight := 0
	if n.left != nil {
		lHeight = n.left.height() + 1
	}
	rHeight := 0
	if n.right != nil {
		rHeight = n.right.height() + 1
	}
	if lHeight > rHeight {
		return lHeight
	}
	return rHeight
}

func (n *Node) String() string {
	parentID := -1
	if n.parent != nil {
		parentID = n.parent.id
	}

	sibling := -1
	if n.parent != nil {
		if n.parent.left != nil && n.parent.left.id != n.id {
			sibling = n.parent.left.id
		} else if n.parent.right != nil && n.parent.right.id != n.id {
			sibling = n.parent.right.id
		}
	}

	degree := 0
	if n.left != nil {
		degree++
	}
	if n.right != nil {
		degree++
	}

	typeS := "internal node"
	if n.parent == nil {
		typeS = "root"
	} else if n.left == nil && n.right == nil {
		typeS = "leaf"
	}

	return fmt.Sprintf("node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s", n.id, parentID, sibling, degree, n.depth(), n.height(), typeS)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	ns := make([]*Node, n)
	for ni := 0; ni < n; ni++ {
		nodeID := nextInt(sc)
		if ns[nodeID] == nil {
			ns[nodeID] = &Node{id: nodeID}
		}

		leftID := nextInt(sc)
		if leftID != -1 {
			if ns[leftID] == nil {
				ns[leftID] = &Node{id: leftID}
			}
			ns[leftID].parent = ns[nodeID]
			ns[nodeID].left = ns[leftID]
		}

		rightID := nextInt(sc)
		if rightID != -1 {
			if ns[rightID] == nil {
				ns[rightID] = &Node{id: rightID}
			}
			ns[rightID].parent = ns[nodeID]
			ns[nodeID].right = ns[rightID]
		}
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
