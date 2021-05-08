package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func sum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

func check(dp [][]int, l, r int) {
	for i := 0; i <= l; i++ {
		for j := 0; j <= r; j++ {
			if (l != i || r != j) && dp[l][r] == dp[i][j] {
				fmt.Println("Yes")
				fmt.Printf("%d", j-i+1)
				for x := i; x <= j; x++ {
					fmt.Printf(" %d", x+1)
				}
				fmt.Println()
				fmt.Printf("%d", r-l+1)
				for x := l; x <= r; x++ {
					fmt.Printf(" %d", x+1)
				}
				fmt.Println()
				os.Exit(0)
			}
		}
	}
}

// FIXME test_17.txtがACしてない
func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt(sc)
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
	}

	for l := 0; l < N; l++ {
		sum := 0
		for r := l; r < N; r++ {
			sum += A[r]
			dp[l][r] = sum % 200
			check(dp, l, r)
		}
	}
	fmt.Println("No")
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
