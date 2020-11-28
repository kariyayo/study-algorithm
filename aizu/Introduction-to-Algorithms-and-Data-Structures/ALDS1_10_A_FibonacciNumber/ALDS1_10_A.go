package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// n=44で2sec以上かかってタイムアウト
func slowFib(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else {
		return slowFib(n-1) + slowFib(n-2)
	}
}

var memo []int

// メモ化再帰で同じ計算を回避
func memoFib(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 1
	} else {
		a := memo[n-1]
		if a == 0 {
			a = memoFib(n - 1)
			memo[n-1] = a
		}
		b := memo[n-2]
		if b == 0 {
			b = memoFib(n - 2)
			memo[n-2] = b
		}
		return a + b
	}
}

var dp []int

// dpでループで小さい方から計算していく
func dpFib(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
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
	memo = make([]int, n)
	dp = make([]int, n+1)
	result := dpFib(n)
	fmt.Println(result)
}
