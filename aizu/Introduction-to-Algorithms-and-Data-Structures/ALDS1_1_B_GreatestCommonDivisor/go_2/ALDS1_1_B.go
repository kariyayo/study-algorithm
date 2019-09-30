package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gcd(x, y int) int {
	if y == 0 {
		return x
	}
	if x >= y {
		m := x % y
		return gcd(y, m)
	}
	m := y % x
	return gcd(x, m)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	x, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	sc.Scan()
	y, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	a := gcd(x, y)
	fmt.Println(a)
}
