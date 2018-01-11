package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	xs [1000]string
	i  int
}

func (s *Stack) push(x string) {
	s.i++
	s.xs[s.i] = x
}

func (s *Stack) pop() string {
	x := s.xs[s.i]
	s.i--
	return x
}

func toInt(s string) int {
	x, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return x
}

func nextStr(sc *bufio.Scanner) string {
	sc.Scan()
	s := sc.Text()
	return s
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	stack := Stack{}

	sc.Scan()
	ss := sc.Text()
	for _, s := range strings.Split(ss, " ") {
		switch s {
		case "*":
			r := toInt(stack.pop())
			l := toInt(stack.pop())
			stack.push(strconv.Itoa(l * r))
		case "+":
			r := toInt(stack.pop())
			l := toInt(stack.pop())
			stack.push(strconv.Itoa(l + r))
		case "-":
			r := toInt(stack.pop())
			l := toInt(stack.pop())
			stack.push(strconv.Itoa(l - r))
		default:
			stack.push(s)
		}
	}
	fmt.Println(stack.pop())
}
