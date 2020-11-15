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

func (node *Node) delete(x int) bool {
	if node == nil {
		return false
	} else if node.key == x {
		if node.left == nil && node.right == nil {
			node.replaceSelfTo(nil)
		} else if node.left != nil && node.right == nil {
			node.replaceSelfTo(node.left)
		} else if node.left == nil && node.right != nil {
			node.replaceSelfTo(node.right)
		} else {
			// 右の部分木の最小値がInorderでの次の値
			node.key = node.right.getMin()
			node.right.delete(node.key)
		}
		return true
	} else if node.key > x {
		return node.left.delete(x)
	} else {
		return node.right.delete(x)
	}
}

func (node *Node) replaceSelfTo(n *Node) {
	// 置き換え対象の親を、自分の親にする
	if n != nil {
		n.p = node.p
	}
	// 自分の親の子を、置き換え対象にする
	if node.p.left != nil && node.p.left.key == node.key {
		node.p.left = n
	} else {
		node.p.right = n
	}
}

func (node *Node) getMin() int {
	n := node
	result := node.key
	for n.left != nil {
		result = n.left.key
		n = n.left
	}
	return result
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

func (t *Tree) delete(x int) bool {
	return t.root.delete(x)
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
		case "delete":
			v := nextInt(sc)
			t.delete(v)
		default:
			panic("not supported. command:${command")
		}
	}
}
