package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solver(xs, ys string) {
	memo = make([][]int, len(xs))
	for p := 0; p < len(xs); p++ {
		memo[p] = make([]int, len(ys))
	}

	ans := 0
	for i := 0; i < len(xs); i++ {
		cnt := count(xs, i, ys, 0)
		debugPrintf("i:%v, cnt:%v\n", i, cnt)
		if cnt > ans {
			ans = cnt
		}
	}
	fmt.Println(ans)
}

var memo [][]int

func count(xs string, xi int, ys string, yi int) int {
	v := memo[xi][yi]
	if v != 0 {
		return v
	}

	m := len(xs) - xi
	n := len(ys) - yi

	if m == 1 && n == 1 {
		if xs[xi] == ys[yi] {
			memo[xi][yi] = 1
		} else {
			memo[xi][yi] = 0
		}
		return memo[xi][yi]
	}

	if m == 1 {
		for i := yi; i < len(ys); i++ {
			if xs[xi] == ys[i] {
				memo[xi][yi] = 1
				return memo[xi][yi]
			}
		}
		memo[xi][yi] = 0
		return memo[xi][yi]
	}

	if n == 1 {
		for i := xi; i < len(xs); i++ {
			if xs[i] == ys[yi] {
				memo[xi][yi] = 1
				return memo[xi][yi]
			}
		}
		memo[xi][yi] = 0
		return memo[xi][yi]
	}

	if xs[xi] == ys[yi] {
		memo[xi][yi] = count(xs, xi+1, ys, yi+1) + 1
		return memo[xi][yi]
	}
	a := count(xs, xi+1, ys, yi)
	b := count(xs, xi, ys, yi+1)
	if a > b {
		memo[xi][yi] = a
		return memo[xi][yi]
	}
	memo[xi][yi] = b
	return memo[xi][yi]
}

func nextStr(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
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
	q := nextInt(sc)

	for i := 0; i < 2*q; i = i + 2 {
		xs := nextStr(sc)
		ys := nextStr(sc)
		solver(xs, ys)
	}
}

func debugPrintf(format string, a ...interface{}) {
	// fmt.Printf(format, a...)
}
