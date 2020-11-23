package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type PriorityQueue struct {
	start int
	len   int
	heap  []int
}

func NewPriorityQueue() *PriorityQueue {
	heap := make([]int, 2000000)
	return &PriorityQueue{len: 0, heap: heap}
}

func (p *PriorityQueue) insert(k int) {
	me := p.len
	parent := (p.len - 1) / 2
	p.heap[p.len] = k

	for p.heap[parent] < k {
		// 親より自分の方が優先度が高いので、自分を上のレベルへ移動
		tmp := p.heap[parent]
		p.heap[parent] = k
		p.heap[me] = tmp
		me = parent
		parent = (parent - 1) / 2
	}
	p.len++
}

func (p *PriorityQueue) extractMax() int {
	result := p.heap[0]
	p.len--
	p.heap[0] = p.heap[p.len]
	p.heap[p.len] = 0
	p.maxHeapify(0)
	return result
}

func (p *PriorityQueue) maxHeapify(i int) {
	n := p.len
	as := p.heap
	l := 2*i + 1
	r := 2*i + 2
	largest := i
	if l < n && as[l] > as[largest] {
		largest = l
	}
	if r < n && as[r] > as[largest] {
		largest = r
	}
	if i != largest {
		// 子の方が優先度が高いので、自分を下のレベルへ移動
		tmp := as[i]
		as[i] = as[largest]
		as[largest] = tmp
		p.maxHeapify(largest)
	}
}

func debugPrintf(format string, a ...interface{}) {
	// fmt.Printf(format, a...)
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
	queue := NewPriorityQueue()
	for {
		command := nextString(sc)
		switch command {
		case "insert":
			k := nextInt(sc)
			queue.insert(k)
		case "extract":
			m := queue.extractMax()
			fmt.Println(m)
		case "end":
			return
		}
	}
}
