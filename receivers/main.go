package main

import "fmt"

// Asociar una funcion a un tipo en particular, similar a una clase - metodos

type Human struct {
    Name string
}

func main() {
    humanA := Human{
        Name: "John",
    }

    humanA.SayHi()
    humanA.SayHiTo(3)
}

// Agregar metodo a Human struct
func (h Human) SayHi() {
    fmt.Println("Hi:", h.Name)
}
func (h Human) SayHiTo(a int) {
    fmt.Println("Hi:", h.Name, a)
}
