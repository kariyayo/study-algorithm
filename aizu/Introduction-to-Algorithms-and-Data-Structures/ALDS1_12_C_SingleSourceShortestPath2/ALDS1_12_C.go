package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	s int
	t int
	w int
}

type Graph [][]*Edge

const INF = 1000000

func dijkstra(d []int, n int, graph Graph) {
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		// usedじゃない頂点の中で最小の頂点を探す
		minV := -1
		minD := INF
		for j := 0; j < n; j++ {
			if !used[j] && d[j] < minD {
				minD = d[j]
				minV = j
			}
		}
		if minV == -1 {
			break
		}

		// `minV`始点の辺に対する緩和
		for _, e := range graph[minV] {
			edge := e
			if d[edge.t] > d[edge.s]+edge.w {
				d[edge.t] = d[edge.s] + edge.w
			}
		}
		used[minV] = true
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
