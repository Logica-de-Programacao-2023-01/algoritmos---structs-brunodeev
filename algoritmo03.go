package main

import "fmt"

func main() {
	triangle := Triangle{
		base:   12.8,
		altura: 36.5,
	}

	fmt.Printf("A área do triângulo é %.2f", calculate(triangle))

}

type Triangle struct {
	base   float64
	altura float64
}

func calculate(triangle Triangle) float64 {
	result := (triangle.base * triangle.altura) / 2
	return result
}
