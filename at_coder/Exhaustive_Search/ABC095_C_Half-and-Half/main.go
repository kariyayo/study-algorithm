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

	A := nextInt(sc)
	B := nextInt(sc)
	C := nextInt(sc)
	X := nextInt(sc)
	Y := nextInt(sc)

	if A+B < 2*C {
		// A, B only
		ans := A*X + B*Y
		fmt.Println(ans)
		return
	}

	ans := 0
	c := 2 * C
	if X > Y {
		ans = ans + c*Y
		if A < c {
			ans = ans + A*(X-Y)
		} else {
			ans = ans + c*(X-Y)
		}
	} else {
		ans = ans + c*X
		if B < c {
			ans = ans + B*(Y-X)
		} else {
			ans = ans + c*(Y-X)
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

func printArray(xs []int) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func debugPrintf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}
