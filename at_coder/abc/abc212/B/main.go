package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)

	ss := nextString(sc)

	s1 := ss[0]
	s2 := ss[1]
	s3 := ss[2]
	s4 := ss[3]

	if s1 == s2 && s2 == s3 && s3 == s4 {
		fmt.Println("Weak")
		return
	}

	x1, _ := strconv.Atoi(string(s1))
	x2, _ := strconv.Atoi(string(s2))
	xx2 := x2 - 1
	if x2 == 0 {
		xx2 = 9
	}
	x3, _ := strconv.Atoi(string(s3))
	xx3 := x3 - 1
	if x3 == 0 {
		xx3 = 9
	}
	x4, _ := strconv.Atoi(string(s4))
	xx4 := x4 - 1
	if x4 == 0 {
		xx4 = 9
	}

	if x1 == xx2 && x2 == xx3 && x3 == xx4 {
		fmt.Println("Weak")
		return
	}

	fmt.Println("Strong")
}

// ----------
// ----------

// Pair

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

// Set

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

// Convert

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

// Combination
//
// nCk
//   combi := pascalTriangle(n)
//   combi[n][k]
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

// Utilities

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

// Scans

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

// Debug print

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
