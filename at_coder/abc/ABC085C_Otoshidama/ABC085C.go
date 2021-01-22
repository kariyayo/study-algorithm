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
	n := nextInt(sc)
	y := nextInt(sc)

	result10000 := -1
	result5000 := -1
	result1000 := -1
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			v := y - 10000*i - 5000*j
			if v < 0 {
				break
			}
			k := v / 1000
			if i+j+k == n {
				result10000 = i
				result5000 = j
				result1000 = k
			}
		}
	}
	fmt.Printf("%d %d %d\n", result10000, result5000, result1000)
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
