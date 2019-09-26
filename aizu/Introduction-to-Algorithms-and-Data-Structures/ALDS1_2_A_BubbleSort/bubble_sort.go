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
		x := as[n-1]
		j := n - 2
		for ; j >= fixed; j-- {
			if as[j] > x {
				as[j+1] = as[j]
				count++
			} else {
				as[j+1] = x
				x = as[j]
			}
		}
		as[j+1] = x
	}
	printArray(as)
	fmt.Println(count)
}

func printArray(as []int) {
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
