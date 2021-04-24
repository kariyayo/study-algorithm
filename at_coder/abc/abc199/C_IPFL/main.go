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
	ss := make([]int, 2*N)
	for i := 0; i < 2*N; i++ {
		ss[i] = i
	}
	S := nextString(sc)
	Q := nextInt(sc)
	flagT := false
	for i := 1; i <= Q; i++ {
		T := nextInt(sc)
		A := nextInt(sc) - 1
		B := nextInt(sc) - 1
		if T == 1 {
			// A文字目とB文字目を入れ替える
			if flagT {
				if A <= N-1 {
					A = A + N
				} else {
					A = A - N
				}
				if B <= N-1 {
					B = B + N
				} else {
					B = B - N
				}
			}
			tmp := ss[A]
			ss[A] = ss[B]
			ss[B] = tmp
		} else {
			if flagT {
				flagT = false
			} else {
				flagT = true
			}
		}
	}
	if flagT {
		// 前半と後半を入れ替える
		s1 := ss[:N]
		s2 := ss[N:]
		ss = append(s2, s1...)
	}
	for _, idx := range ss {
		fmt.Print(string(S[idx]))
	}
	fmt.Println()
}

// ----------
type Pair []int

func pair(first, second int) Pair {
	return []int{first, second}
}

func (p Pair) first() int {
	return p[0]
}
func (p Pair) second() int {
	return p[1]
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
