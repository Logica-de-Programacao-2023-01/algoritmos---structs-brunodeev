package main

import (
	"fmt"
)

func main() {
	person := Person{
		name: "Bruno",
		age:  18,
		address: Address{
			street: "Rua Governador",
			number: 8,
			city:   "Planaltina",
			state:  "GO",
		},
	}

	fmt.Print(address(person).address.street)
	fmt.Print(" ", address(person).address.number)
	fmt.Print(" ", address(person).address.city)
	fmt.Print("-", address(person).address.state)
}

type Address struct {
	street string
	number int
	city   string
	state  string
}

type Person struct {
	name    string
	age     int
	address Address
}

func address(person Person) Person {
	return person
}
