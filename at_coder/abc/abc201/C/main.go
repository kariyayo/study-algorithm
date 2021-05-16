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

	N := 10
	s := nextString(sc)

	requires := ""
	options := ""
	for i := 0; i < N; i++ {
		s := string(s[i])
		if s == "o" {
			requires = requires + fmt.Sprintf("%d", i)
		} else if s == "?" {
			options = options + fmt.Sprintf("%d", i)
		}
	}
	if len(requires) > 4 {
		fmt.Println(0)
		return
	}

	ans := 0

	for i := 0; i <= 9999; i++ {
		s := fmt.Sprintf("%04d", i)
		usedRequires := NewStringSet()
		okCnt := 0
		for j := 0; j < 4; j++ {
			c := string(s[j])
			if strings.Contains(requires, c) {
				usedRequires.append(c)
				okCnt++
			}
			if strings.Contains(options, c) {
				okCnt++
			}
		}
		if okCnt == 4 && len(usedRequires) == len(requires) {
			ans++
		}
	}

	fmt.Println(ans)
}

func uniqCount(s string) int {
	set := map[byte]struct{}{}
	for i := 0; i < len(s); i++ {
		set[s[i]] = struct{}{}
	}
	return len(set)
}

// ----------
type StringSet map[string]struct{}

func NewStringSet() StringSet {
	return map[string]struct{}{}
}

func (set StringSet) append(s string) {
	set[s] = struct{}{}
}

type Pair struct {
	First  string
	Second int
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

func abs(a int) int {
	return int(math.Abs(float64(a)))
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
	fmt.Fprintln(os.Stderr, strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintTable(table [][]int) {
	for i := 0; i < len(table); i++ {
		debugPrintArray(table[i])
	}
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, format, a...)
}
