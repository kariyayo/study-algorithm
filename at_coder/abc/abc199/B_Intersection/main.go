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
	buf := make([]byte, 0)
	sc.Buffer(buf, 1010000)
	sc.Split(bufio.ScanWords)

	N := nextInt(sc)
	as := make([]int, N)
	bs := make([]int, N)
	for i := 0; i < N; i++ {
		as[i] = nextInt(sc)
	}
	for i := 0; i < N; i++ {
		bs[i] = nextInt(sc)
	}
	xmap := make(map[int]int, 1000)
	for i := 0; i < N; i++ {
		a := as[i]
		b := bs[i]
		if a > b {
			continue
		}
		for x := a; x <= b; x++ {
			_, ok := xmap[x]
			if ok {
				xmap[x] = xmap[x] + 1
			} else {
				xmap[x] = 1
			}
		}
	}
	ans := 0
	for _, cnt := range xmap {
		if cnt == N {
			ans++
		}
	}
	fmt.Println(ans)
}

// ----------
type Pair []int

func pair(first, second int) Pair {
	return []int{first, second}
}

func (p Pair) first() int {
	return p[0]
}
func (p Pair) second() int {
	return p[1]
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
