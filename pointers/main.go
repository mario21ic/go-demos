package main

import (
    "fmt"
)

func main() {
    var myVar string = "original message"
    // name myVar
    // type string
    // address 0x00???
    // value "original message"

    // Puntero: variable que almacena la direccion en memoria de un valor
    var myPtr *string = &myVar


    fmt.Printf("myVar value: %v\n", myVar)
    fmt.Printf("myVar direcc: %v\n", &myVar)
    fmt.Printf("myPtr value: %v\n", myPtr)
    fmt.Printf("myPtr direcc: %v\n", &myPtr)

    // Dereferenciar el puntero *nombrePuntero
    fmt.Printf("Valor al que apunta myPtr (*ptr): %v\n", *myPtr)

    // Modificar el valor al que apunta el puntero
    *myPtr = "abc"
    fmt.Println()
    fmt.Printf("Valor de myVar: %v\n", myVar)
}
