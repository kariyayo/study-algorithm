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

func bellmanFord(d []int, n int, graph Graph) {
	for i := 0; i < n; i++ {
		isUpdate := false
		// 各辺に対して緩和
		for v := 0; v < n; v++ {
			if d[v] == INF {
				continue
			}
			for _, e := range graph[v] {
				edge := e
				if d[edge.t] > d[edge.s]+edge.w {
					d[edge.t] = d[edge.s] + edge.w
					isUpdate = true
				}
			}
		}
		if !isUpdate {
			break
		}
		if i == n && isUpdate {
			fmt.Println("NEGATIVE CYCLE!")
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

	bellmanFord(d, n, graph)
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
