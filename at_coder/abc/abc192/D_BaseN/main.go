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

	X := nextString(sc)
	M := nextUint64(sc)

	if len(X) == 1 {
		x, err := strconv.Atoi(X)
		if err != nil {
			panic(err)
		}
		if uint64(x) <= M {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
		return
	}

	ss := strings.Split(X, "")

	d := 0
	for _, s := range ss {
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		d = max(d, x)
	}

	// 何進数までOKかを二分探索する
	ac := uint64(d)
	wa := M + 1
	for (wa - ac) > 1 {
		var wj uint64 = (ac + wa) / 2
		var v uint64 = 0
		for _, c := range ss {
			if v > M/wj { // オーバーフロー対策
				v = M + 1
			} else {
				x, err := strconv.ParseUint(c, 10, 64)
				if err != nil {
					panic(err)
				}
				v = v*wj + x
			}
		}
		if v <= M {
			ac = wj
		} else {
			wa = wj
		}
	}
	// d+1進数からac進数がOKなので、ac-d種類が答え
	fmt.Println(ac - uint64(d))
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
