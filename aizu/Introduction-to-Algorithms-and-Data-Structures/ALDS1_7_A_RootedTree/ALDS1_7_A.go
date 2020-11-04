package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tree struct {
	nodes []*Node
}

func NewTree(size int) *Tree {
	nodes := make([]*Node, size)
	return &Tree{nodes}
}

func (t *Tree) append(id int, children []int) {
	var node *Node
	if t.nodes[id] == nil {
		node = &Node{id: id, children: children}
	} else {
		node = t.nodes[id]
		node.children = children
	}
	t.nodes[id] = node
	for _, childID := range children {
		if t.nodes[childID] == nil {
			child := &Node{id: childID}
			t.nodes[childID] = child
		}
		t.nodes[childID].parent = node
	}
}

func (t *Tree) Println() {
	for i := 0; i < len(t.nodes); i++ {
		t.nodes[i].Println()
	}
}

type Node struct {
	id       int
	parent   *Node
	children []int
}

func (n *Node) Println() {
	var d = 0
	p := n.parent
	for p != nil {
		d++
		p = p.parent
	}
	var parentID = -1
	if n.parent != nil {
		parentID = n.parent.id
	}

	cntChildren := len(n.children)

	t := "internal node"
	if parentID == -1 {
		t = "root"
	} else if cntChildren == 0 {
		t = "leaf"
	}

	fmt.Printf("node %d: parent = %d, depth = %d, %s, ", n.id, parentID, d, t)

	fmt.Print("[")
	i := 0
	for i < cntChildren {
		fmt.Printf("%d", n.children[i])
		i++
		if i != cntChildren {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	tree := NewTree(n)
	for i := 0; i < n; i++ {
		id := nextInt(sc)
		childrenSize := nextInt(sc)
		children := make([]int, childrenSize)
		for j := 0; j < childrenSize; j++ {
			children[j] = nextInt(sc)
		}
		tree.append(id, children)
	}
	tree.Println()
}
