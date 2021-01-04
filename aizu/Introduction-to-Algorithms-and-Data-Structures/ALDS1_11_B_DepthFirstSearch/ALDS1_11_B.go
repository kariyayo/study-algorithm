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
	f  int
}

func (v *Vertex) visited() bool {
	return v.d != -1
}

func search(g []*Vertex, t int, vertex *Vertex) int {
	vertex.d = t + 1
	f := vertex.d
	for _, vID := range vertex.vs {
		next := g[vID]
		if next.visited() {
			continue
		}
		res := search(g, f, next)
		if res > f {
			f = res
		}
	}
	vertex.f = f + 1
	return vertex.f
}

func solver(g []*Vertex) {
	t := 0
	for _, vertex := range g[1:] {
		if vertex.visited() {
			continue
		}
		t = search(g, t, vertex)
	}
	for _, vertex := range g[1:] {
		fmt.Printf("%d %d %d\n", vertex.id, vertex.d, vertex.f)
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
		g[u] = &Vertex{u, k, vs, -1, -1}
	}
	solver(g)
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
