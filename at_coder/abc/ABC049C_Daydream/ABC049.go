package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	s = strings.ReplaceAll(s, "eraser", "")
	s = strings.ReplaceAll(s, "erase", "")
	s = strings.ReplaceAll(s, "dreamer", "")
	s = strings.ReplaceAll(s, "dream", "")
	if len(s) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

// ----------

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
