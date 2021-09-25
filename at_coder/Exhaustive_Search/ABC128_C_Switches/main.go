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
	M := nextInt(sc)
	l := make([][]int, M)
	for i := 0; i < M; i++ {
		k := nextInt(sc)
		ss := make([]int, k)
		for j := 0; j < k; j++ {
			ss[j] = nextInt(sc) - 1
		}
		l[i] = ss
	}
	ps := make([]int, M)
	for i := 0; i < M; i++ {
		ps[i] = nextInt(sc)
	}

	ans := 0
	for bit := 0; bit < (1 << N); bit++ {
		lighted := 0
		for i, ss := range l {
			// 電球ごとのループ
			onSwitch := 0
			for _, s := range ss {
				if bit&(1<<s) != 0 {
					onSwitch++
				}
			}
			if onSwitch%2 == ps[i] {
				lighted++
			}
		}
		if lighted == M {
			// 全ての電球が点灯
			ans++
		}
	}
	fmt.Println(ans)
}

// ----------

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
