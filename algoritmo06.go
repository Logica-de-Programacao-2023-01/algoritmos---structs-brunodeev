package main

import "fmt"

type Funcionário struct {
	Nome    string
	Salario float64
	Idade   int
}

func salaryChange(f Funcionário) float64 {

	newSalary := f.Salario + f.Salario*0.10

	return newSalary
}

func timeService(f Funcionário) int {
	tempoS := f.Idade - 18
	if f.Idade < 18 {
		return 0
	}
	return tempoS
}

func main() {
	s := salaryChange(Funcionário{
		Nome:    "Joana",
		Salario: 40000,
		Idade:   37,
	})
	fmt.Println("O novo salário é: ", s)

	t := timeService(Funcionário{
		Nome:  "Joana",
		Idade: 37,
	})
	fmt.Println("O tempo de serviço do funcionário é: ", t, "anos")
}
