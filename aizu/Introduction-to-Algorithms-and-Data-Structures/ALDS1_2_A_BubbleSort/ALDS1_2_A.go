package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func bubbleSort(n int, as []int) {
	count := 0
	for fixed := 0; fixed < n; fixed++ {
		target := as[n-1]
		for j := n - 2; j >= fixed; j-- {
			if as[j] > target {
				as[j+1] = as[j]
				count++
			} else {
				as[j+1] = target
				target = as[j]
			}
		}
		as[fixed] = target
	}
	printArr(as)
	fmt.Println(count)
}

func printArr(as []int) {
	fmt.Println(strings.Trim(fmt.Sprint(as), "[]"))
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

	bubbleSort(n, as)
}
