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
	n := nextInt(sc)
	x := 0
	y := 0
	before := 0
	ok := true
	for i := 0; i < n; i++ {
		t := nextInt(sc)
		xt := nextInt(sc)
		yt := nextInt(sc)
		distance := int(math.Abs(float64(xt-x))) + int(math.Abs(float64(yt-y)))
		dt := t - before
		if distance-dt > 0 || (dt-distance)%2 != 0 {
			ok = false
			break
		}
		x = xt
		y = yt
		before = t
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
