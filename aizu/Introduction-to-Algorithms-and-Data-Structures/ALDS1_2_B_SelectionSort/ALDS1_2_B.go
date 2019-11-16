package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SelectionSort(n uint8, as []uint8) {
	cnt := 0
	for i := uint8(0); i < n; i++ {
		minj := i
		for j := i; j < n; j++ {
			if as[j] < as[minj] {
				minj = j
			}
		}
		if minj > i {
			tmp := as[minj]
			as[minj] = as[i]
			as[i] = tmp
			cnt++
		}
	}
	printArray(as)
	fmt.Println(cnt)
}

func printArray(as []uint8) {
	fmt.Println(strings.Trim(fmt.Sprint(as), "[]"))
}

func nextInt(sc *bufio.Scanner) uint8 {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return uint8(n)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	as := make([]uint8, n)
	for i := uint8(0); i < n; i++ {
		x := nextInt(sc)
		as[i] = x
	}
	SelectionSort(n, as)
}
