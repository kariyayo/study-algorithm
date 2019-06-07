package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func isPrime(x int) bool {
	if x == 0 {
		return true
	} else if x == 1 {
		return true
	} else if x == 2 {
		return true
	} else if x%2 == 0 {
		return false
	} else {
		y := math.Sqrt(float64(x))
		for i := 3; float64(i) <= y; i = i + 2 {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
}

func prime(xs []int) {
	count := 0
	for _, x := range xs {
		if isPrime(x) {
			count = count + 1
		}
	}
	fmt.Println(count)
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	s := sc.Text()
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	n := nextInt(sc)

	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = nextInt(sc)
	}
	prime(xs)
}
