package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func search(v int64, xs []int64) int {
	if len(xs) == 0 {
		return -1
	}
	m := len(xs) / 2
	x := xs[m]
	if v < x {
		return search(v, xs[:m])
	} else {
		return search(v, xs[m:])
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)
	Q := nextInt(sc)

	as := make([]int64, N+1)
	cs := make([]int64, N+1)
	for i := 1; i <= N; i++ {
		as[i] = nextInt64(sc)
		cs[i] = as[i] - 1 - as[i-1] + cs[i-1]
	}
	for i := 0; i < Q; i++ {
		K := nextInt64(sc)
		if K > cs[N] {
			fmt.Println(as[N] + K - cs[N])
		} else {
			wa := 0
			ac := N
			for (ac - wa) > 1 {
				m := (wa + ac) / 2
				if K <= cs[m] {
					ac = m
				} else {
					wa = m
				}
			}
			fmt.Println(as[wa] + (K - cs[wa]))
		}
	}
}

// ----------
type Pair struct {
	First  int
	Second int
}

type SIPair struct {
	First  string
	Second int
}

type LLPair struct {
	First  int64
	Second int64
}

type IntSet map[int]struct{}

func NewIntSet() IntSet {
	return map[int]struct{}{}
}
func (set IntSet) append(x int) {
	set[x] = struct{}{}
}

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return map[string]struct{}{}
}
func (set StringSet) append(s string) {
	set[s] = struct{}{}
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func toInt64(s string) int64 {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
}

func pascalTriangle(n int) [][]int {
	pt := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		pt[i] = make([]int, i+1)
		pt[i][0] = 1
		pt[i][i] = 1
		for j := 1; j < i; j++ {
			pt[i][j] = pt[i-1][j-1] + pt[i-1][j]
		}
	}
	return pt
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

func nextInt64(sc *bufio.Scanner) int64 {
	sc.Scan()
	n, err := strconv.ParseInt(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func printArray(xs []int) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintArray(xs []int) {
	fmt.Fprintln(os.Stderr, strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintTable(table [][]int) {
	for i := 0; i < len(table); i++ {
		debugPrintArray(table[i])
	}
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
