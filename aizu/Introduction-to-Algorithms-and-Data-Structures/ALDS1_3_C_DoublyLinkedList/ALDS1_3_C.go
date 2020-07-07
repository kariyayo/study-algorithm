package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func (n Node) String() string {
	prev := 0
	if n.prev != nil {
		prev = n.prev.value
	}
	next := 0
	if n.next != nil {
		next = n.next.value
	}
	return fmt.Sprintf("prev:%d, value:%d, next:%d", prev, n.value, next)
}

type DoublyLinkedList struct {
	start *Node
	last  *Node
	size  int
}

func (l *DoublyLinkedList) insert(x int) {
	node := &Node{x, nil, l.start}
	if l.size == 0 {
		l.start = node
		l.last = node
	} else if l.size != 0 {
		node.next.prev = node
		l.start = node
	}
	l.size++
}

func (l *DoublyLinkedList) delete(x int) {
	if l.size == 0 {
		return
	} else if l.size == 1 {
		if l.start.value == x {
			l.start = nil
			l.last = nil
			l.size = 0
		}
		return
	} else {
		for node := l.start; node != nil; node = node.next {
			if node.value == x {
				if node.prev != nil {
					node.prev.next = node.next
				} else {
					l.start = node.next
				}
				if node.next != nil {
					node.next.prev = node.prev
				} else {
					l.last = node.prev
				}
				l.size--
				return
			}
		}
	}
}

func (l *DoublyLinkedList) deleteFirst() {
	if l.size == 1 {
		l.start = nil
		l.last = nil
		l.size = 0
	} else if l.size > 1 {
		l.start.next.prev = nil
		l.start = l.start.next
		l.size--
	}
}

func (l *DoublyLinkedList) deleteLast() {
	if l.size == 1 {
		l.start = nil
		l.last = nil
		l.size = 0
	} else if l.size > 1 {
		l.last.prev.next = nil
		l.last = l.last.prev
		l.size--
	}
}

func (l DoublyLinkedList) String() string {
	node := l.start
	xs := make([]string, l.size)
	for i := 0; i < l.size; i++ {
		xs[i] = fmt.Sprintf("%d", node.value)
		node = node.next
	}
	return strings.Join(xs, " ")
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n := toInt(sc.Text())

	l := DoublyLinkedList{}

	for i := 0; i < n; i++ {
		sc.Scan()
		command := sc.Text()
		switch command {
		case "insert":
			sc.Scan()
			v := toInt(sc.Text())
			l.insert(v)
		case "delete":
			sc.Scan()
			v := toInt(sc.Text())
			l.delete(v)
		case "deleteFirst":
			l.deleteFirst()
		case "deleteLast":
			l.deleteLast()
		}
	}

	fmt.Println(l)
}
