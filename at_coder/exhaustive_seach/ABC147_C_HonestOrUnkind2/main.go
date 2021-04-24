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
	as := make([][]Pair, N)
	for i := 0; i < N; i++ {
		A := nextInt(sc)
		ps := make([]Pair, A)
		for j := 0; j < A; j++ {
			x := nextInt(sc)
			y := nextInt(sc)
			ps[j] = pair(x, y)
		}
		as[i] = ps
	}

	ans := 0
	for bit := 0; bit < (1 << N); bit++ {
		cnt := 0
		for i := 0; i < N; i++ {
			if bit&(1<<i) != 0 {
				cnt++
				// as[i]は全て正しい
				ps := as[i]
				for _, p := range ps {
					if p.second() == 1 && bit&(1<<(p.first()-1)) == 0 {
						// 矛盾
						cnt = 0
						break
					} else if p.second() == 0 && bit&(1<<(p.first()-1)) != 0 {
						// 矛盾
						cnt = 0
						break
					}
				}
				if cnt == 0 {
					break
				}
			}
		}
		if cnt > 0 {
			ans = max(ans, cnt)
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
