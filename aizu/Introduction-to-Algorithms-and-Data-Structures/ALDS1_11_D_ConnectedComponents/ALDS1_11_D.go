package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vertex struct {
	id  int
	vs  []int
	gid int // 部分グラフID
}

func dfs(g []*Vertex, vertex *Vertex, gid int) {
	vertex.gid = gid
	for _, nextVID := range vertex.vs {
		next := g[nextVID]
		if next.gid != 0 {
			continue
		}
		dfs(g, next, gid)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	m := nextInt(sc)
	g := make([]*Vertex, n)
	for i := 0; i < n; i++ {
		g[i] = &Vertex{id: i, vs: []int{}, gid: 0}
	}
	for i := 0; i < m; i++ {
		id1 := nextInt(sc)
		id2 := nextInt(sc)
		g[id1].vs = append(g[id1].vs, id2)
		g[id2].vs = append(g[id2].vs, id1)
	}

	gid := 1
	for i := 0; i < n; i++ {
		if g[i].gid == 0 {
			dfs(g, g[i], gid)
		}
		gid++
	}

	q := nextInt(sc)
	for i := 0; i < q; i++ {
		s := nextInt(sc)
		t := nextInt(sc)
		if g[s].gid == g[t].gid {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
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

func printArray(xs []int) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
