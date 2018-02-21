package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type point struct {
	x float64
	y float64
}

var th = math.Pi * 60.0 / 180.0

func kock(l point, r point, x int) {
	if x == 0 {
		fmt.Printf("%f %f\n", l.x, l.y)
	} else {
		p1 := point{
			x: (2.0*l.x + 1.0*r.x) / 3.0,
			y: (2.0*l.y + 1.0*r.y) / 3.0,
		}
		p2 := point{
			x: (1.0*l.x + 2.0*r.x) / 3.0,
			y: (1.0*l.y + 2.0*r.y) / 3.0,
		}
		m := point{
			x: (p2.x-p1.x)*math.Cos(th) - (p2.y-p1.y)*math.Sin(th) + p1.x,
			y: (p2.x-p1.x)*math.Sin(th) + (p2.y-p1.y)*math.Cos(th) + p1.y,
		}
		kock(l, p1, x-1)
		kock(p1, m, x-1)
		kock(m, p2, x-1)
		kock(p2, r, x-1)
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	l := point{x: 0.0, y: 0.0}
	r := point{x: 100.0, y: 0.0}
	kock(l, r, n)
	fmt.Printf("%f %f\n", r.x, r.y)
}
