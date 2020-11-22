package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	key      int
	priority int
	parent   *Node
	left     *Node
	right    *Node
}

func (node *Node) accept(visitor Visitor) {
	visitor.visitNode(node)
}

func (node *Node) insert(key, priority int) *Node {
	if node == nil {
		return &Node{key: key, priority: priority}
	} else if node.key == key {
		return node
	} else if node.key > key {
		// 左部分木にinsert
		node.left = node.left.insert(key, priority)
		if node.left.priority > node.priority {
			// heap条件を満たすために右回転
			node = node.rightRotate()
		}
		return node
	} else {
		// 右部分木にinsert
		node.right = node.right.insert(key, priority)
		if node.right.priority > node.priority {
			// heap条件を満たすために左回転
			node = node.leftRotate()
		}
		return node
	}
}

func (node *Node) rightRotate() *Node {
	s := node.left
	node.left = s.right
	s.right = node
	s.parent = node.parent
	node.parent = s
	return s
}

func (node *Node) leftRotate() *Node {
	s := node.right
	node.right = s.left
	s.left = node
	s.parent = node.parent
	node.parent = s
	return s
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

func (node *Node) delete(x int) *Node {
	if node == nil {
		return nil
	} else if node.key > x {
		node.left = node.left.delete(x)
		return node
	} else if node.key < x {
		node.right = node.right.delete(x)
		return node
	}

	// 以降、node.key == x の場合
	if node.left == nil && node.right == nil {
		return nil
	}
	var n *Node
	if node.left != nil && node.right == nil {
		// 左の子を持ち上げる
		n = node.rightRotate()
	} else if node.left == nil && node.right != nil {
		// 右の子を持ち上げる
		n = node.leftRotate()
	} else {
		// priorityが大きい方を持ち上げる
		if node.left.priority > node.right.priority {
			n = node.rightRotate()
		} else {
			n = node.leftRotate()
		}
	}
	return n.delete(x)
}

func (node *Node) replaceSelfTo(n *Node) {
	// 置き換え対象の親を、自分の親にする
	if n != nil {
		n.parent = node.parent
	}
	// 自分の親の子を、置き換え対象にする
	if node.parent.left != nil && node.parent.left.key == node.key {
		node.parent.left = n
	} else {
		node.parent.right = n
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

func (t *Tree) insert(key int, priority int) {
	node := t.root.insert(key, priority)
	if node.parent == nil {
		t.root = node
	}
}

func (t *Tree) find(x int) bool {
	if t.root == nil {
		return false
	}
	return t.root.find(x)
}

func (t *Tree) delete(x int) {
	t.root.delete(x)
	node := t.root
	for node.parent != nil {
		node = node.parent
	}
	t.root = node
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
			p := nextInt(sc)
			t.insert(v, p)
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
