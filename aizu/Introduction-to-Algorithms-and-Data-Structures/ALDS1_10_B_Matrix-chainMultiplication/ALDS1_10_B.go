package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type matrix [2]int

func (m matrix) rows() int {
	return m[0]
}

func (m matrix) cols() int {
	return m[1]
}

func newMatrix(rows, cols int) matrix {
	return [2]int{rows, cols}
}

var ms []matrix

var memo [][]int

func solver(start, last int) int {
	m := memo[start][last]
	if m > 0 {
		return m
	}

	min := 1000001
	if last-start <= 0 {
		memo[start][last] = 0
		return 0
	} else if last-start == 1 {
		result := ms[start].rows() * ms[start].cols() * ms[last].cols()
		memo[start][last] = result
		return result
	}

	for k := 1; start+k <= last; k++ {
		cnt1 := solver(start, start+k-1)
		cnt2 := solver(start+k, last)
		cnt3 := ms[start].rows() * ms[start+k].rows() * ms[last].cols()
		cnt := cnt1 + cnt2 + cnt3
		if cnt < min {
			min = cnt
		}
	}

	if min == 1000001 {
		memo[start][last] = 0
		return 0
	}
	memo[start][last] = min
	return min
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
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	ms = make([]matrix, n)
	for i := 0; i < n; i++ {
		rows := nextInt(sc)
		cols := nextInt(sc)
		ms[i] = newMatrix(rows, cols)
	}
	memo = make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
	}

	result := solver(0, n-1)
	fmt.Println(result)
}

func debugPrintf(format string, a ...interface{}) {
	// fmt.Printf(format, a...)
}
