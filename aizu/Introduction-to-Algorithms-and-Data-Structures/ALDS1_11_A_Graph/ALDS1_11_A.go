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
}

func solver(g []*Vertex) {
	n := len(g)
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for _, vertex := range g {
		u := vertex.id
		for _, v := range vertex.vs {
			ans[u-1][v-1] = 1
		}
	}
	for i := 0; i < n; i++ {
		printArray(ans[i])
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
	g := make([]*Vertex, n)
	for i := 0; i < n; i++ {
		u := nextInt(sc)
		k := nextInt(sc)
		vs := make([]int, k)
		for j := 0; j < k; j++ {
			vs[j] = nextInt(sc)
		}
		g[i] = &Vertex{u, k, vs}
	}
	solver(g)
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
