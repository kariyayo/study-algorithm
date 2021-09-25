package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func factorial(n int) int {
	result := 1
	for i := n; i > 0; i-- {
		result = result * i
	}
	return result
}

func iter(v string, vs []string) [][]string {
	if len(vs) == 1 {
		result := []string{v, vs[0]}
		return [][]string{result}
	}
	children := [][]string{}
	for i := 0; i < len(vs); i++ {
		vv := vs[i]
		left := vs[:i]
		right := []string{}
		if i < len(vs) {
			right = vs[i+1:]
		}
		others := make([]string, len(left))
		copy(others, left)
		others = append(others, right...)
		children = append(children, iter(vv, others)...)
	}
	result := make([][]string, len(children))
	for i := 0; i < len(children); i++ {
		result[i] = []string{v}
		result[i] = append(result[i], children[i]...)
	}
	return result
}

func permutation(vs []string) [][]string {
	result := [][]string{}
	for i := 0; i < len(vs); i++ {
		v := vs[i]
		left := vs[:i]
		right := []string{}
		if i < len(vs) {
			right = vs[i+1:]
		}
		others := make([]string, len(left))
		copy(others, left)
		others = append(others, right...)
		result = append(result, iter(v, others)...)
	}
	return result
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)
	vs := make([]string, N)
	P := make([]string, N)
	for i := 0; i < N; i++ {
		s := nextString(sc)
		P[i] = s
		vs[i] = s
	}
	Q := make([]string, N)
	for i := 0; i < N; i++ {
		Q[i] = nextString(sc)
	}

	// 順列を作ってから並べ替えると遅いので事前に並べ替えておく
	sort.Slice(vs, func(i, j int) bool {
		return vs[i] < vs[j]
	})

	// vsの順列
	patterns := permutation(vs)

	a := 0
	b := 0
	for i := 0; i < len(patterns); i++ {
		pattern := patterns[i]
		if fmt.Sprintf("%v", P) == fmt.Sprintf("%v", pattern) {
			a = i + 1
		}
		if fmt.Sprintf("%v", Q) == fmt.Sprintf("%v", pattern) {
			b = i + 1
		}
	}
	fmt.Println(abs(a - b))
}

// ----------
type Pair struct {
	First  int
	Second int
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
	fmt.Fprintf(os.Stderr, "\n---\n%v\n", strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
