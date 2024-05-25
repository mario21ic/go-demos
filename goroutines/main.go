package main

import (
    "fmt"
    "sync"
)

// Usaremos waitgroup para esperar a que termine la goroutine
func Hola(wg *sync.WaitGroup) {
    fmt.Println("hola")
    wg.Done()
}
func Hello(wg *sync.WaitGroup) {
    fmt.Println("hello")
    wg.Done()
}

func main() {
    fmt.Println("## Starting")

    var wg sync.WaitGroup
    wg.Add(2) // nro de goroutines
    go Hola(&wg)
    go Hello(&wg)
    wg.Wait() // esperamos que terminen

    fmt.Println("## Finished")
}
