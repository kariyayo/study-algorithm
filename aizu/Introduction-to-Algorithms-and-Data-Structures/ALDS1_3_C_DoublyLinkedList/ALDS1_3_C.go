package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	Value int
	Next  int
	Prev  int
}

type DoublyLinkedList struct {
	ns   [2000001]*Node
	head int
	tail int
	size int
}

func (l *DoublyLinkedList) insert(x int) {
	if l.head == cap(l.ns)-1 {
		l.head = 0
	} else {
		l.head++
	}
	if l.size == 0 {
		l.ns[l.head] = &Node{
			Value: x,
			Next:  -1,
			Prev:  -1,
		}
		l.tail = l.head
	} else {
		var next int
		if l.head == 0 {
			next = cap(l.ns) - 1
		} else {
			next = l.head - 1
		}
		l.ns[l.head] = &Node{
			Value: x,
			Next:  next,
			Prev:  -1,
		}
		l.ns[next].Prev = l.head
	}
	l.size++
}

func (l *DoublyLinkedList) delete(x int) {
	n := l.ns[l.head]
	for x != n.Value && n.Next != -1 {
		n = l.ns[n.Next]
	}
	if x == n.Value {
		// 削除
		if n.Prev > -1 {
			l.ns[n.Prev].Next = n.Next
		}
		if n.Next > -1 {
			l.ns[n.Next].Prev = n.Prev
		}
	}
	// 先頭要素が削除された場合
	if n.Prev == -1 {
		l.head = n.Next
	}
	// 最後尾要素が削除された場合
	if n.Next == -1 {
		l.tail = n.Prev
	}
	l.size--
}

func (l *DoublyLinkedList) deleteFirst() {
	n := l.ns[l.head]
	if n.Next > -1 {
		l.ns[n.Next].Prev = n.Prev
	}
	l.head = n.Next
	l.size--
}

func (l *DoublyLinkedList) deleteLast() {
	n := l.ns[l.tail]
	l.tail = n.Prev
	if n.Prev > -1 {
		l.ns[n.Prev].Next = -1
	}
	l.size--
}

func (l *DoublyLinkedList) println() {
	n := l.ns[l.head]
	fmt.Print(n.Value)
	for n.Next != -1 {
		n = l.ns[n.Next]
		fmt.Print(" ")
		fmt.Print(n.Value)
	}
	fmt.Println("")
}

func nextStr(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return s
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

	var l DoublyLinkedList
	n := nextInt(sc)
	for i := 0; i < n; i++ {
		s := nextStr(sc)
		switch s {
		case "insert":
			x := nextInt(sc)
			l.insert(x)
		case "delete":
			x := nextInt(sc)
			l.delete(x)
		case "deleteFirst":
			l.deleteFirst()
		case "deleteLast":
			l.deleteLast()
		}
	}
	l.println()
}
