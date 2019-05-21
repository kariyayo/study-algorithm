package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func insertionSort(n int, as []int) {
	printArray(as)
	for i := 1; i < n; i++ {
		v := as[i]
		j := i - 1
		for ; j >= 0; j-- {
			if v > as[j] {
				break
			}
			as[j+1] = as[j]
		}
		as[j+1] = v
		printArray(as)
	}
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
	insertionSort(n, as)
}
