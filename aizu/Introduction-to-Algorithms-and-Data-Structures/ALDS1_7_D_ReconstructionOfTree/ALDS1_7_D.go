package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var i = 0

func solve(n int, ans []int, preorder, inorder []int) []int {
	inorderIndex := -1
	for j := 0; j < len(inorder); j++ {
		v := inorder[j]
		if preorder[i] == v {
			inorderIndex = j
		}
	}

	leftTree := inorder[:inorderIndex]
	rightTree := inorder[inorderIndex+1:]

	switch len(leftTree) {
	case 0:
		// NOP
	case 1:
		i++
		ans = append(ans, leftTree[0])
	default:
		i++
		ans = solve(n, ans, preorder, leftTree)
	}

	switch len(rightTree) {
	case 0:
		// NOP
	case 1:
		i++
		ans = append(ans, rightTree[0])
	default:
		i++
		ans = solve(n, ans, preorder, rightTree)
	}

	ans = append(ans, inorder[inorderIndex])
	return ans
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return n
}

func printArray(xs []int) {
	fmt.Println(strings.Trim(fmt.Sprint(xs), "[]"))
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	n := nextInt(sc)
	preorder := make([]int, n)
	inorder := make([]int, n)
	for i := 0; i < n; i++ {
		v := nextInt(sc)
		preorder[i] = v
	}
	for i := 0; i < n; i++ {
		v := nextInt(sc)
		inorder[i] = v
	}

	ans := []int{}
	ans = solve(n, ans, preorder, inorder)
	printArray(ans)
}
