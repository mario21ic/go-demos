package main

import "fmt"

// Nuevo tipo llamado day
type day int

func main() {
	var d day = 26
	printDay(d)

	// Descomentar para probar que no permite el tipo por no ser day
	// var i int = 26
	// printDay(i)
}

func printDay(d day) {
	fmt.Printf("Dias: %d\n", d)
}
