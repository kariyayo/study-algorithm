package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func search(n int, as []int, q int, bs []int) {
	count := 0
	for _, b := range bs {
		flag := true
		for i := 0; i < n && flag; i++ {
			if b == as[i] {
				flag = false
				count++
			}
		}
	}
	fmt.Println(count)
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
	var as []int
	for i := 0; i < n; i++ {
		as = append(as, nextInt(sc))
	}
	q := nextInt(sc)
	var bs []int
	for i := 0; i < q; i++ {
		bs = append(bs, nextInt(sc))
	}
	search(n, as, q, bs)
}
