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
	h := nextInt(sc)
	w := nextInt(sc)
	table := make([][]int, h) // 0:白, 1:黒
	cnt := 0
	for i := 0; i < h; i++ {
		table[i] = make([]int, w)
		cs := nextString(sc)
		for j, c := range cs {
			if c == '#' {
				table[i][j] = 1
				cnt++
			}
		}
	}
	ans := 0
	for i := 0; i < h-1; i++ {
		for j := 0; j < w-1; j++ {
			p1 := table[i][j]
			p2 := table[i+1][j]
			p3 := table[i][j+1]
			p4 := table[i+1][j+1]
			sum := p1 + p2 + p3 + p4
			if sum == 1 || sum == 3 {
				ans++
			}
		}
	}
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
