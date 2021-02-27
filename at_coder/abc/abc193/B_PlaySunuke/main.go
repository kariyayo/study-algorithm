package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const _INF = 1000000001

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)
	ans := _INF
	for i := 0; i < N; i++ {
		A := nextInt(sc)
		P := nextInt(sc)
		X := nextInt(sc)
		if X-((60*A-30)/60) > 1 {
			// 買える
			ans = min(ans, P)
		}
	}
	if ans == _INF {
		ans = -1
	}
	fmt.Println(ans)
}

// ----------

func min(a, b int) int {
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

func debugPrintf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
