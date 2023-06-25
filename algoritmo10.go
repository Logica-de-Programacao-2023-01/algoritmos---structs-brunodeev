package main

import "fmt"

type Filme struct {
	Título     string
	Diretor    string
	Ano        int
	Avaliações []int
}

func addRating([]int) []int {
	numeros := []int{8, 6, 9, 9}
	numeros = append(numeros, 10)

	return numeros
}

func deleteRating([]int) []int {
	nums := []int{8, 6, 9, 9}
	nums = nums[:len(nums)-1]

	return nums
}

func averageRating([]int) float64 {
	numeros := []int{8, 6, 9, 9}
	media := (numeros[0] + numeros[1] + numeros[2] + numeros[3]) / 4

	return float64(media)
}

func main() {
	r := Filme{
		Título:     "Zootopia",
		Diretor:    "Rich Moore",
		Ano:        2016,
		Avaliações: []int{8, 6, 9, 9},
	}
	fmt.Println(r)

	f := averageRating([]int{8, 6, 9, 9})
	fmt.Println("A média das avaliações é: ", f)
}
