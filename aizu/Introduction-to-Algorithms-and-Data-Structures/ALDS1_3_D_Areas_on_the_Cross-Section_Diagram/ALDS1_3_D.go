package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	xs []int
}

func (s *Stack) push(x int) {
	s.xs = append(s.xs, x)
}

func (s *Stack) pop() (int, bool) {
	if len(s.xs) == 0 {
		return 0, false
	}
	x := s.xs[len(s.xs)-1]
	s.xs = s.xs[:len(s.xs)-1]
	return x, true
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	s := sc.Text()

	stack := &Stack{}
	areas := make(map[int]int)
	totalArea := 0
	last := 0
	for currentStep, c := range s {
		switch c {
		case '\\':
			stack.push(currentStep)
		case '/':
			startStep, ok := stack.pop()
			if ok {
				area := currentStep - startStep
				totalArea += area
				areas[startStep] += area
				for i := startStep + 1; i < currentStep; i++ {
					a, ok := areas[i]
					if ok {
						areas[startStep] += a
						delete(areas, i)
					}
				}
				last = startStep
			}
		}
	}
	fmt.Println(totalArea)
	fmt.Printf("%d", len(areas))
	for i := 0; i <= last; i++ {
		x, ok := areas[i]
		if ok {
			fmt.Printf(" %d", x)
		}
	}
	fmt.Println()
}
