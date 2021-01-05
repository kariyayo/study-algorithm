package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Queue struct {
	xs   [100001]int
	head int
	tail int
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) enq(x int) {
	q.xs[q.tail] = x
	if q.tail >= len(q.xs)-1 {
		q.tail = 0
	} else {
		q.tail++
	}
}

func (q *Queue) deq() int {
	x := q.xs[q.head]
	if q.head >= len(q.xs)-1 {
		q.head = 0
	} else {
		q.head++
	}
	return x
}

func (q *Queue) empty() bool {
	return q.head == q.tail
}

type Vertex struct {
	id int
	k  int
	vs []int
	d  int
}

func bfs(g []*Vertex, start *Vertex) {
	queue := NewQueue()

	start.d = 0
	queue.enq(start.id)

	for !queue.empty() {
		vID := queue.deq()
		vertex := g[vID]
		for _, nextVID := range vertex.vs {
			next := g[nextVID]
			if next.d != -1 {
				continue
			}
			next.d = g[vID].d + 1
			queue.enq(next.id)
		}
	}
}

func solver(g []*Vertex) {
	start := g[1]
	bfs(g, start)
	for _, vertex := range g[1:] {
		fmt.Printf("%d %d\n", vertex.id, vertex.d)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	g := make([]*Vertex, n+1)
	for i := 0; i < n; i++ {
		u := nextInt(sc)
		k := nextInt(sc)
		vs := make([]int, k)
		for j := 0; j < k; j++ {
			vs[j] = nextInt(sc)
		}
		g[u] = &Vertex{u, k, vs, -1}
	}
	solver(g)
}

//----------

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
	fmt.Printf(format, a...)
}
