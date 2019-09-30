package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func gcd(x, y int) {
	for {
		if x == 0 || y == 0 {
			fmt.Println(0)
			return
		} else if x == 1 || y == 1 {
			fmt.Println(1)
			return
		} else if x > y && x%y == 0 {
			fmt.Println(y)
			return
		} else if y > x && y%x == 0 {
			fmt.Println(y)
			return
		} else if x == y {
			fmt.Println(x)
			return
		}
		tmp := y
		y = x % y
		x = tmp
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	s := sc.Text()
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	sc.Scan()
	s = sc.Text()
	y, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	gcd(x, y)
}
