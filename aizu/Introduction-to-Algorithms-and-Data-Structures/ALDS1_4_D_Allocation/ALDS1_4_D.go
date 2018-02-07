package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 最大積載量がpの時、k台のトラックで何個詰めるかを返します
func check(n int, k int, ws []int, p int64) int {
	i := 0
	for j := 0; j < k; j++ {
		s := 0
		for int64(s+ws[i]) <= p {
			s = s + ws[i]
			i++
			if i == n {
				return n
			}
		}
	}
	return i
}

func resolve(n int, k int, ws []int) int64 {
	var left int64
	left = 0
	var right int64
	right = 10000 * 100000
	var mid int64
	for (right - left) > 1 {
		mid = int64((right + left) / 2)
		v := check(n, k, ws, mid)
		if v >= n {
			// 積載量midで十分に詰める
			right = mid
		} else {
			// 積載量midでは足りない
			left = mid
		}
	}
	return right
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
	k := nextInt(sc)
	var ws []int
	for i := 0; i < n; i++ {
		ws = append(ws, nextInt(sc))
	}
	p := resolve(n, k, ws)
	fmt.Println(p)
}
