package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func g1(x int) int {
	s := fmt.Sprintf("%d", x)
	ss := strings.Split(s, "")
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	result, err := strconv.Atoi(strings.Join(ss, ""))
	if err != nil {
		panic(err)
	}
	return result
}

func g2(x int) int {
	s := fmt.Sprintf("%d", x)
	ss := strings.Split(s, "")
	sort.Strings(ss)
	result, err := strconv.Atoi(strings.Join(ss, ""))
	if err != nil {
		panic(err)
	}
	return result
}

func f(x int) int {
	return g1(x) - g2(x)
}

func solver(n int, k int) int {
	ans := n
	for i := 0; i < k; i++ {
		ans = f(ans)
	}
	return ans
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	k := nextInt(sc)
	ans := solver(n, k)
	fmt.Println(ans)
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
