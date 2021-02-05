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
	var lNode *Node
	if left != -1 {
		if t.nodes[left] == nil {
			lNode = &Node{id: left, parent: -1}
			t.nodes[left] = lNode
		} else {
			lNode = t.nodes[left]
		}
		lNode.parent = id
	}
	var rNode *Node
	if right != -1 {
		if t.nodes[right] == nil {
			rNode = &Node{id: right, parent: -1}
			t.nodes[right] = rNode
		} else {
			rNode = t.nodes[right]
		}
		rNode.parent = id
	}
	var node *Node
	if t.nodes[id] == nil {
		node = &Node{id: id, parent: -1, left: lNode, right: rNode}
		t.nodes[id] = node
	} else {
		node = t.nodes[id]
	}
	node.left = lNode
	node.right = rNode
}

func (t *Tree) root() int {
	node := t.nodes[0]
	for node.parent != -1 {
		node = t.nodes[node.parent]
	}
	return node.id
}

func (t *Tree) Walk(visitor TreeVisitor) Doc {
	return t.nodes[t.root()].accept(visitor)
}

type Node struct {
	id     int
	parent int
	left   *Node
	right  *Node
}

func (n *Node) accept(visitor TreeVisitor) Doc {
	if n == nil {
		return visitor.VisitLeaf()
	}
	return visitor.VisitNode(n)
}

type Doc string

type TreeVisitor interface {
	VisitLeaf() Doc
	VisitNode(node *Node) Doc
}

type preOrderVisitor struct{}

func (v *preOrderVisitor) VisitLeaf() Doc {
	return ""
}
func (v *preOrderVisitor) VisitNode(node *Node) Doc {
	s := fmt.Sprintf(" %d", node.id)
	l := node.left.accept(v)
	r := node.right.accept(v)
	return Doc(s) + l + r
}

type inOrderVisitor struct{}

func (v *inOrderVisitor) VisitLeaf() Doc {
	return ""
}
func (v *inOrderVisitor) VisitNode(node *Node) Doc {
	l := node.left.accept(v)
	s := fmt.Sprintf(" %d", node.id)
	r := node.right.accept(v)
	return l + Doc(s) + r
}

type postOrderVisitor struct{}

func (v *postOrderVisitor) VisitLeaf() Doc {
	return ""
}
func (v *postOrderVisitor) VisitNode(node *Node) Doc {
	l := node.left.accept(v)
	r := node.right.accept(v)
	s := fmt.Sprintf(" %d", node.id)
	return l + r + Doc(s)
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
	fmt.Println(tree.Walk(&preOrderVisitor{}))
	fmt.Println("Inorder")
	fmt.Println(tree.Walk(&inOrderVisitor{}))
	fmt.Println("Postorder")
	fmt.Println(tree.Walk(&postOrderVisitor{}))
}
