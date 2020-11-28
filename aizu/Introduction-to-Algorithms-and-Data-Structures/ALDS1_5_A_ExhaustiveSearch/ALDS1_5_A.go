package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func solve(as []int, m int, i int) bool {
	if m == 0 {
		return true
	}
	if i > len(as)-1 {
		return false
	}
	res := solve(as, m-as[i], i+1)
	if res {
		return true
	}
	res = solve(as, m, i+1)
	if res {
		return true
	}
	return false
}

func solver(as []int, ms []int) {
	for _, m := range ms {
		result := solve(as, m, 0)
		if result {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)
	as := make([]int, n)
	for i := 0; i < n; i++ {
		as[i] = nextInt(sc)
	}
	q := nextInt(sc)
	ms := make([]int, q)
	for i := 0; i < q; i++ {
		ms[i] = nextInt(sc)
	}

	solver(as, ms)
}
