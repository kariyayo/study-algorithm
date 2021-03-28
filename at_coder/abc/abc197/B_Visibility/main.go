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

	H := nextInt(sc)
	W := nextInt(sc)
	x := nextInt(sc) - 1
	y := nextInt(sc) - 1

	row := ""
	col := make([]byte, H)
	for i := 0; i < H; i++ {
		S := nextString(sc)
		if x == i {
			row = S
		}
		col[i] = S[y]
	}

	ans := 1
	// up
	for i := x - 1; i >= 0; i-- {
		if string(col[i]) == "#" {
			break
		}
		ans++
	}
	// down
	for i := x + 1; i < H; i++ {
		if string(col[i]) == "#" {
			break
		}
		ans++
	}
	// left
	for i := y - 1; i >= 0; i-- {
		if string(row[i]) == "#" {
			break
		}
		ans++
	}
	// right
	for i := y + 1; i < W; i++ {
		if string(row[i]) == "#" {
			break
		}
		ans++
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
