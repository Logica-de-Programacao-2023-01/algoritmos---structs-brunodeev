package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	Especie string
	Idade   int
	Som     string
}

func soundChange(a Animal) string {
	som := a.Som
	novoSom := strings.ReplaceAll(a.Som, som, "miau")

	return novoSom
}

func main() {
	l := Animal{
		Especie: "Felis catus",
		Idade:   2,
	}

	an := soundChange(Animal{
		Especie: "Felis catus",
		Idade:   2,
		Som:     "meow",
	})
	fmt.Println(l, an)

}
