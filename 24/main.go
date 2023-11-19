package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func (p Point) DistanceTo(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}
func main() {
	a := Point{1, 2}
	b := Point{1, 3}
	fmt.Println(a.DistanceTo(b))
}
