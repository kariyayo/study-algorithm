package main

import (
	"bufio"
	"bytes"
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

var preorder bytes.Buffer
var inorder bytes.Buffer
var postorder bytes.Buffer

func walk(node *Node) {
	if node == nil {
		return
	}
	fmt.Fprintf(&preorder, " %d", node.id)
	walk(node.left)
	fmt.Fprintf(&inorder, " %d", node.id)
	walk(node.right)
	fmt.Fprintf(&postorder, " %d", node.id)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	ns := make([]*Node, n)
	for i := 0; i < n; i++ {
		id := nextInt(sc)
		node := ns[id]
		if node == nil {
			node = &Node{id: id}
			ns[id] = node
		}

		leftID := nextInt(sc)
		if leftID != -1 {
			leftNode := ns[leftID]
			if leftNode == nil {
				leftNode = &Node{id: leftID}
				ns[leftID] = leftNode
			}
			node.left = leftNode
			leftNode.parent = node
		}

		rightID := nextInt(sc)
		if rightID != -1 {
			rightNode := ns[rightID]
			if rightNode == nil {
				rightNode = &Node{id: rightID}
				ns[rightID] = rightNode
			}
			node.right = rightNode
			rightNode.parent = node
		}
	}

	root := ns[0]
	for root.parent != nil {
		root = root.parent
	}

	walk(root)

	fmt.Println("Preorder")
	fmt.Println(preorder.String())

	fmt.Println("Inorder")
	fmt.Println(inorder.String())

	fmt.Println("Postorder")
	fmt.Println(postorder.String())
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
