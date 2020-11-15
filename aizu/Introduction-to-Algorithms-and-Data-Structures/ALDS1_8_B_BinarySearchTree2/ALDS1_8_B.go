package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key   int
	p     *Node
	left  *Node
	right *Node
}

func (node *Node) accept(visitor Visitor) {
	visitor.visitNode(node)
}

func (node *Node) find(x int) bool {
	if node.key == x {
		return true
	} else if node.key > x {
		if node.left == nil {
			return false
		}
		return node.left.find(x)
	} else {
		if node.right == nil {
			return false
		}
		return node.right.find(x)
	}
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(z *Node) {
	var p *Node
	m := t.root
	for m != nil {
		p = m
		if m.key > z.key {
			m = m.left
		} else {
			m = m.right
		}
	}
	z.p = p

	if p == nil {
		t.root = z
	} else if p.key > z.key {
		p.left = z
	} else {
		p.right = z
	}
}

func (t *Tree) find(x int) bool {
	if t.root == nil {
		return false
	}
	return t.root.find(x)
}

type Visitor interface {
	visitNode(node *Node)
}

type InorderVisitor struct {
	ss []string
}

func (i *InorderVisitor) visitNode(node *Node) {
	if node == nil {
		return
	}
	node.left.accept(i)
	i.ss = append(i.ss, fmt.Sprintf("%d", node.key))
	node.right.accept(i)
}

func (i *InorderVisitor) String() string {
	return strings.Join(i.ss, " ")
}

type PreorderVisitor struct {
	ss []string
}

func (p *PreorderVisitor) visitNode(node *Node) {
	if node == nil {
		return
	}
	p.ss = append(p.ss, fmt.Sprintf("%d", node.key))
	node.left.accept(p)
	node.right.accept(p)
}

func (p *PreorderVisitor) String() string {
	return strings.Join(p.ss, " ")
}

func nextString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
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
	t := &Tree{}
	for i := 0; i < n; i++ {
		command := nextString(sc)
		switch command {
		case "insert":
			v := nextInt(sc)
			t.insert(&Node{key: v})
		case "print":
			if t.root != nil {
				inorderVisitor := &InorderVisitor{}
				t.root.accept(inorderVisitor)
				fmt.Printf(" %s\n", inorderVisitor)
				preorderVisitor := &PreorderVisitor{}
				t.root.accept(preorderVisitor)
				fmt.Printf(" %s\n", preorderVisitor)
			}
		case "find":
			v := nextInt(sc)
			if t.find(v) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		default:
			panic("not supported. command:${command")
		}
	}
}
