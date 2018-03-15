package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partition(xs []int, p int, r int) int {
	x := xs[r]
	as := xs[:r]
	i := p - 1 // 前パーティションの最後の要素のポインタ
	j := p
	for j < len(as) {
		target := as[j]
		if target <= x {
			i++
			as[j] = as[i]
			as[i] = target
		}
		j++
	}
	xs[r] = xs[i+1]
	xs[i+1] = x
	return i + 1
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func printArray(xs []int, q int) {
	ss := make([]string, len(xs))
	for i, x := range xs {
		if i == q {
			ss[i] = "[" + strconv.Itoa(x) + "]"
		} else {
			ss[i] = strconv.Itoa(x)
		}
	}
	fmt.Println(strings.Join(ss, " "))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = nextInt(sc)
	}
	q := partition(xs, 0, n-1)
	printArray(xs, q)
}
