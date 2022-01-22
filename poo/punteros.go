package main

// Nota: no hay operaciones aritmeticas en punteros

func main() {
	var i *int
	var y *int

	c := i + y

	a := []int{1, 2, 3}
	p := &a[0]
	p++
}
