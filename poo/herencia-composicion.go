package main

import (
	"fmt"
)

func main() {
	g := gopher{animal{"Mario's Gopher"}, 4}
	g.Print()
}

type animal struct {
	name string
}

func (a animal) Print() {
	fmt.Printf("animalPrint: I'm %s and I'm a %T", a.name, a)
}

type gopher struct {
	animal
	legs int
}

// Descomentar para ver herencia
// func (g gopher) Print() {
// 	fmt.Printf("gopherPrint: I'm %s and I'm a %T", g.name, g)
// }
