package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const INF = math.MaxInt32

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)
	M := nextInt(sc)
	S := make([]string, N)
	T := make([]string, M)
	for i := 0; i < N; i++ {
		S[i] = nextString(sc)
	}
	for i := 0; i < M; i++ {
		T[i] = nextString(sc)
	}

	i := 0
	j := 0
	for i < N {
		s := S[i]
		t := T[j]
		if s == t {
			fmt.Println("Yes")
			i++
			j++
		} else {
			fmt.Println("No")
			i++
		}
	}
}

// ----------
// ----------

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

// key以上の要素のうちMINである要素の添字を返す
// ref. https://gist.github.com/bati11/1fa106711d17ddea6ab39c43e989fea3
func lowerBound(xs []int, key int) int {
	if len(xs) == 0 {
		return 0
	}
	l, r := 0, len(xs)-1
	if xs[r] < key {
		// key以上の要素が存在しない
		return r + 1
	}
	for r-l > 1 {
		mid := l + (r-l)/2
		v := xs[mid]
		if key <= v {
			r = mid
		} else {
			l = mid
		}
	}
	if xs[l] >= key {
		return l
	}
	return r
}

// keyより大きな要素のうちMINである要素の添字を返す
// ref. https://gist.github.com/bati11/1fa106711d17ddea6ab39c43e989fea3
func upperBound(xs []int, key int) int {
	if len(xs) == 0 {
		return 0
	}
	l, r := 0, len(xs)-1
	if xs[r] <= key {
		// keyより大きい要素が存在しない
		return r + 1
	}
	for r-l > 1 {
		mid := l + (r-l)/2
		v := xs[mid]
		if key < v {
			r = mid
		} else {
			l = mid
		}
	}
	if xs[l] > key {
		return l
	}
	return r
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

// Union-Find
type UnionFind struct {
	parents []int
	size    []int
}

func (u *UnionFind) root(x int) int {
	if u.parents[x] == -1 {
		return x
	}
	p := u.root(u.parents[x])
	// 親を根で置き換えてしまう
	//   経路圧縮で効率化 O(N) -> O(logN)
	u.parents[x] = p
	return p
}

func (u *UnionFind) Unite(x, y int) {
	xRoot := u.root(x)
	yRoot := u.root(y)
	// 既に同じグループの場合は何もしない
	if xRoot == yRoot {
		return
	}
	// サイズが小さい方のグループを大きい方のグループに統合する
	//   union by size による効率化 O(logN) -> O(α(N))
	if u.size[xRoot] > u.size[yRoot] {
		u.parents[yRoot] = xRoot
		u.size[xRoot] = u.size[xRoot] + u.size[yRoot]
	} else {
		u.parents[xRoot] = yRoot
		u.size[yRoot] = u.size[yRoot] + u.size[xRoot]
	}
}

func (u *UnionFind) Same(x, y int) bool {
	return u.root(x) == u.root(y)
}

func NewUnionFind(n int) *UnionFind {
	parents := make([]int, n)
	for i := 0; i < n; i++ {
		parents[i] = -1
	}
	// sizeの大小比較ができれば良いので特に初期化していないが、
	// 値自体に意味がある場合は 1 で初期化する
	size := make([]int, n)
	return &UnionFind{parents: parents, size: size}
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

func makeInts(n int, defaultValue int) []int {
	result := make([]int, n)
	if defaultValue != 0 {
		for i := 0; i < n; i++ {
			result[i] = defaultValue
		}
	}
	return result
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
	if a < 0 {
		return -1 * a
	}
	return a
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
