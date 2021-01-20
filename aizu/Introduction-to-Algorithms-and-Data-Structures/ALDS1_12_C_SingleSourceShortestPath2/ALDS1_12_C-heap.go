package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	k   int
	vID int
}

type PriorityQueue struct {
	start int
	len   int
	heap  []Node
}

func NewPriorityQueue() *PriorityQueue {
	heap := make([]Node, 2000000)
	return &PriorityQueue{len: 0, heap: heap}
}

func (p *PriorityQueue) insert(node Node) {
	me := p.len
	parent := (p.len - 1) / 2
	p.heap[p.len] = node

	for p.heap[parent].k > node.k {
		// 親より自分の方が優先度が低いので、自分を上のレベルへ移動
		tmp := p.heap[parent]
		p.heap[parent] = node
		p.heap[me] = tmp
		me = parent
		parent = (parent - 1) / 2
	}
	p.len++
}

func (p *PriorityQueue) extractMin() Node {
	result := p.heap[0]
	p.len--
	p.heap[0] = p.heap[p.len]
	p.heap[p.len] = Node{}
	p.minHeapify(0)
	return result
}

func (p *PriorityQueue) minHeapify(i int) {
	n := p.len
	as := p.heap
	l := 2*i + 1
	r := 2*i + 2
	largest := i
	if l < n && as[l].k < as[largest].k {
		largest = l
	}
	if r < n && as[r].k < as[largest].k {
		largest = r
	}
	if i != largest {
		// 子の方が優先度が低いので、自分を下のレベルへ移動
		tmp := as[i]
		as[i] = as[largest]
		as[largest] = tmp
		p.minHeapify(largest)
	}
}

type Edge struct {
	s int
	t int
	w int
}

type Graph [][]*Edge

const INF = 1000000

func dijkstra(d []int, n int, graph Graph) {
	heap := NewPriorityQueue()
	heap.insert(Node{k: 0, vID: 0})
	for heap.len != 0 {
		// usedじゃない頂点の中で最小の頂点`minV`を探す
		// このminVを探す部分がheapを使うことで疎グラフの場合に効率よくなる
		minNode := heap.extractMin()
		minV := minNode.vID
		minD := minNode.k
		// 緩和で特定の頂点に対するdが変わった時に、ヒープにinsertするだけで古いdを持ったNodeを削除していない
		// そのため、最小値を持たないNodeは無視する
		if minD > d[minV] {
			continue
		}
		if minV == -1 {
			break
		}

		// `minV`始点の辺に対する緩和
		for _, e := range graph[minV] {
			edge := e
			if d[edge.t] > d[edge.s]+edge.w {
				d[edge.t] = d[edge.s] + edge.w
				// 更新されたらheapに追加（古いNodeの削除はしない。実装を簡単にするため）
				heap.insert(Node{k: d[edge.t], vID: edge.t})
			}
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d %d\n", i, d[i])
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	d := make([]int, n)
	graph := make([][]*Edge, n)
	for i := 0; i < n; i++ {
		u := nextInt(sc)
		d[u] = INF
		k := nextInt(sc)
		es := make([]*Edge, k)
		for j := 0; j < k; j++ {
			v := nextInt(sc)
			c := nextInt(sc)
			edge := &Edge{s: u, t: v, w: c}
			es[j] = edge
		}
		graph[u] = es
	}
	d[0] = 0

	dijkstra(d, n, graph)
}

// ----------

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

func printArray(xs []*Edge) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
