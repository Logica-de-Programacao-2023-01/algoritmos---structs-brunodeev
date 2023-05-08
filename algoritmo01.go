package main

import (
	"fmt"
	"math"
)

func main() {
	circle := Circle{
		radius: 9.75,
	}

	fmt.Printf("%.1f", area(circle))
}

type Circle struct {
	radius float64
}

func area(circle Circle) float64 {
	calculate := math.Pi * (circle.radius * circle.radius)
	return calculate
}
