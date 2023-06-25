package main

import "fmt"

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func adicionarNota(a Aluno) []float64 {
	n := a.Notas
	n = append(a.Notas, 7)

	return n
}

func removerNota(a Aluno) []float64 {
	d := a.Notas
	d = a.Notas[:len(a.Notas)-1]

	return d
}

func mediaNotas([]float64) float64 {
	notas := []float64{9, 8, 10, 7}
	avg := (notas[0] + notas[1] + notas[2] + notas[3]) / float64(len(notas))

	return avg
}

func main() {
	a := Aluno{
		Nome:  "Carolina",
		Idade: 18,
		Notas: []float64{9, 8, 10, 7},
	}
	fmt.Println(a)

	b := mediaNotas([]float64{9, 8, 10, 7})

	fmt.Println("A média das notas desse aluno é: ", b)
}
