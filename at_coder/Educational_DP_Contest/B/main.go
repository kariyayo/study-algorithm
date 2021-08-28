package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const INF = 1010101010

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)
	K := nextInt(sc)
	hs := make([]int, N)
	dp := make([]int, N)
	for i := 0; i < N; i++ {
		hs[i] = nextInt(sc)
		dp[i] = INF
	}
	dp[0] = 0

	for i := 0; i < N; i++ {
		for k := 1; k <= K; k++ {
			// i番目からk個ジャンプする
			if i+k < N {
				dp[i+k] = min(dp[i+k], dp[i]+abs(hs[i+k]-hs[i]))
			}
		}
	}
	fmt.Println(dp[N-1])
}

// ----------
// ----------

// table
func createTable(h int, w int) [][]int {
	table := make([][]int, h)
	for i := 0; i < h; i++ {
		table[i] = make([]int, w)
	}
	return table
}

// Slice

func removeByte(slice []byte, s int) []byte {
	return append(slice[:s], slice[s+1:]...)
}

func removeInt(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

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
func (set IntSet) put(x int) {
	set[x] = struct{}{}
}

type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return map[string]struct{}{}
}
func (set StringSet) append(s string) {
	set[s] = struct{}{}
}

// Priority Queue
type Item struct {
	value    int64
	priority int64
	index    int
}
type PriorityQueue []*Item

func (q PriorityQueue) Len() int { return len(q) }

func (q PriorityQueue) Less(i, j int) bool { return q[i].priority > q[j].priority }

func (q PriorityQueue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *PriorityQueue) Push(x interface{}) {
	n := len(*q)
	item := x.(*Item)
	item.index = n
	*q = append(*q, item)
}

func (q *PriorityQueue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*q = old[0 : n-1]
	return item
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

// Permutations
// nPn
func permutations(arr []byte) [][]byte {
	var helper func([]byte, int)
	res := [][]byte{}

	helper = func(arr []byte, n int) {
		if n == 1 {
			tmp := make([]byte, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
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
	fmt.Fprintf(os.Stdout, format, a...)
}
