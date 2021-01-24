package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	as := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = nextInt(sc)
	}
	maxSum := 0
	for l := 0; l < n; l++ {
		x := 100000
		for r := l; r < n; r++ {
			x = min(x, as[r])
			sum := x * (r - l + 1)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	fmt.Println(maxSum)
}

// ----------

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
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

func printArray(xs []int) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
