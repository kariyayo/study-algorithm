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

func (t *Tree) append(id int, left int, right int) {
	var node *Node
	if t.nodes[id] == nil {
		node = &Node{id: id, left: left, right: right}
		t.nodes[id] = node
	} else {
		node = t.nodes[id]
		node.left = left
		node.right = right
	}
	for _, childID := range []int{left, right} {
		if childID == -1 {
			continue
		}
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

func (t *Tree) root() int {
	node := t.nodes[0]
	for node.parent != nil {
		node = node.parent
	}
	return node.id
}

func (t *Tree) Preorder() {
	t.preorder(t.root())
}

func (t *Tree) preorder(id int) {
	if id < 0 {
		return
	}
	node := t.nodes[id]
	fmt.Printf(" %d", node.id)
	t.preorder(node.left)
	t.preorder(node.right)
}

func (t *Tree) Ineorder() {
	t.inorder(t.root())
}

func (t *Tree) inorder(id int) {
	if id < 0 {
		return
	}
	node := t.nodes[id]
	t.inorder(node.left)
	fmt.Printf(" %d", node.id)
	t.inorder(node.right)
}

func (t *Tree) Postorder() {
	t.postorder(t.root())
}

func (t *Tree) postorder(id int) {
	if id < 0 {
		return
	}
	node := t.nodes[id]
	t.postorder(node.left)
	t.postorder(node.right)
	fmt.Printf(" %d", node.id)
}

type Node struct {
	id     int
	parent *Node
	left   int
	right  int
}

func (n *Node) Println() {
	parentID := -1
	if n.parent != nil {
		parentID = n.parent.id
	}

	sibling := -1
	if n.parent != nil {
		if n.parent.left != -1 && n.parent.left != n.id {
			sibling = n.parent.left
		} else if n.parent.right != -1 && n.parent.right != n.id {
			sibling = n.parent.right
		}
	}

	degree := 0
	if n.left != -1 {
		degree++
	}
	if n.right != -1 {
		degree++
	}

	depth := 0
	p := n.parent
	for p != nil {
		depth++
		p = p.parent
	}

	nodeType := "internal node"
	if n.parent == nil {
		nodeType = "root"
	} else if degree == 0 {
		nodeType = "leaf"
	}

	fmt.Printf("node %d: parent = %d, sibling = %d, degree = %d, depth = %d, %s\n", n.id, parentID, sibling, degree, depth, nodeType)
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
		left := nextInt(sc)
		right := nextInt(sc)
		tree.append(id, left, right)
	}
	fmt.Println("Preorder")
	tree.Preorder()
	fmt.Println("")
	fmt.Println("Inorder")
	tree.Ineorder()
	fmt.Println("")
	fmt.Println("Postorder")
	tree.Postorder()
	fmt.Println("")
}
