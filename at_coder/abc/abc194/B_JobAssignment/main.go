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
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)

	A1 := []int{0, 100000}
	A2 := []int{0, 100000}
	B1 := []int{0, 100000}
	B2 := []int{0, 100000}
	for i := 1; i <= N; i++ {
		a := nextInt(sc)
		if A1[1] > a {
			A2[0] = A1[0]
			A2[1] = A1[1]
			A1[0] = i
			A1[1] = a
		} else if A2[1] > a {
			A2[0] = i
			A2[1] = a
		}
		b := nextInt(sc)
		if B1[1] > b {
			B2[0] = B1[0]
			B2[1] = B1[1]
			B1[0] = i
			B1[1] = b
		} else if B2[1] > b {
			B2[0] = i
			B2[1] = b
		}
	}
	ans := 300000
	if A1[0] != B1[0] {
		ans = max(A1[1], B1[1])
	} else {
		ans = A1[1] + B1[1]
		if A2[1] < B2[1] {
			x := max(A2[1], B1[1])
			ans = min(ans, x)
		} else {
			x := max(A1[1], B2[1])
			ans = min(ans, x)
		}
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
