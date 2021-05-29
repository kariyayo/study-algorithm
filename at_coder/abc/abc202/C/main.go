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

	N := nextInt(sc)
	A := make([]int, N)
	aMap := map[int]int{} // {1:1, 2:2}
	for i := 0; i < N; i++ {
		a := nextInt(sc)
		A[i] = a
		_, ok := aMap[a]
		if ok {
			aMap[a]++
		} else {
			aMap[a] = 1
		}
	}
	B := make([]int, N)
	matchedB := map[int]int{} // { 1:1, 2:2 }
	for i := 0; i < N; i++ {
		b := nextInt(sc)
		B[i] = b
		cnt, ok := aMap[b]
		if ok {
			_, ok2 := matchedB[i]
			if ok2 {
				matchedB[i] = matchedB[i] + cnt
			} else {
				matchedB[i] = cnt
			}
		}
	}
	C := make([]int, N)
	ans := 0
	for i := 0; i < N; i++ {
		c := nextInt(sc)
		C[i] = c
		cnt, ok := matchedB[c-1]
		if ok {
			ans = ans + cnt
		}
	}

	fmt.Println(ans)
}

func uniqCount(s string) int {
	set := map[byte]struct{}{}
	for i := 0; i < len(s); i++ {
		set[s[i]] = struct{}{}
	}
	return len(set)
}

// ----------
type Pair struct {
	First  string
	Second int
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

func toInt64(s string) uint64 {
	x, err := strconv.ParseUint(s, 10, 64)
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

func minUint64(a, b uint64) uint64 {
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

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func abs(a int) int {
	return int(math.Abs(float64(a)))
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

func nextUint64(sc *bufio.Scanner) uint64 {
	sc.Scan()
	n, err := strconv.ParseUint(sc.Text(), 10, 64)
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
