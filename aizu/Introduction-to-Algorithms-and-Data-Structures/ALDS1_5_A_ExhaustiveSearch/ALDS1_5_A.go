package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var as []int

func solve(m int, i int) string {
	if m == 0 {
		return "yes"
	}
	if i > len(as)-1 {
		return "no"
	}
	res := solve(m-as[i], i+1)
	if res == "yes" {
		return "yes"
	}
	res = solve(m, i+1)
	if res == "yes" {
		return "yes"
	}
	return "no"
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
	for n > 0 {
		as = append(as, nextInt(sc))
		n--
	}
	q := nextInt(sc)
	var ms []int
	for q > 0 {
		ms = append(ms, nextInt(sc))
		q--
	}

	for _, m := range ms {
		result := solve(m, 0)
		fmt.Println(result)
	}
}
