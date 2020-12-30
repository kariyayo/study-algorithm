package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var maxLen = 1001

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solver(n int, x, y string) {
	dp := make([][]int, len(x)+1)
	for i := range dp {
		dp[i] = make([]int, len(y)+1)
	}

	for i := 1; i <= len(x); i++ {
		for j := 1; j <= len(y); j++ {
			if i == 1 && j == 1 {
				if x[i-1] == y[j-1] {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
			} else {
				if x[i-1] == y[j-1] {
					dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]+1)
				} else {
					dp[i][j] = max(dp[i][j-1], dp[i-1][j])
				}
			}
		}
	}
	fmt.Println(dp[len(x)][len(y)])
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

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	n := nextInt(sc)
	for i := 0; i < n; i++ {
		x := nextString(sc)
		y := nextString(sc)
		solver(n, x, y)
	}
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
