package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	id int
	k  int
	vs []int
	d  int
}

func visit(g []*Vertex, vertex *Vertex, d int) {
	// 既にセット済みの距離の方が近い場合はスキップ
	if vertex.d != -1 && vertex.d <= d {
		return
	}
	vertex.d = d
	distance := d + 1
	for _, vID := range vertex.vs {
		next := g[vID]
		visit(g, next, distance)
	}
}

func solver(g []*Vertex) {
	start := g[1]
	visit(g, start, 0)
	for _, vertex := range g[1:] {
		fmt.Printf("%d %d\n", vertex.id, vertex.d)
	}
}

func printArray(xs []int) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
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

func debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
