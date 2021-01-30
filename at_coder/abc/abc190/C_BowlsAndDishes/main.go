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
	m := nextInt(sc)
	goals := make([][]int, m)
	for i := 0; i < m; i++ {
		a := nextInt(sc)
		b := nextInt(sc)
		pair := make([]int, 2)
		pair[0] = a
		pair[1] = b
		goals[i] = pair
	}
	k := nextInt(sc)
	parsons := make([][]int, k)
	for i := 0; i < k; i++ {
		c := nextInt(sc)
		d := nextInt(sc)
		parson := make([]int, 2)
		parson[0] = c
		parson[1] = d
		parsons[i] = parson
	}

	answer := 0
	for bit := 0; bit < (1 << k); bit++ {
		a := make([]bool, n+1)
		for i := 0; i < k; i++ {
			p := parsons[i]
			if (bit>>i)&1 != 0 {
				// Cを使う
				a[p[0]] = true
			} else {
				// Dを使う
				a[p[1]] = true
			}
		}
		tmpAnswer := 0
		for _, goal := range goals {
			matchCount := 0
			if a[goal[0]] {
				matchCount++
			}
			if a[goal[1]] {
				matchCount++
			}
			if matchCount >= 2 {
				tmpAnswer++
			}
		}
		if tmpAnswer > answer {
			answer = tmpAnswer
		}
	}
	fmt.Println(answer)
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
