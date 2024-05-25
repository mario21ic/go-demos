package main

import "fmt"

// Los canales ponen Lock, por eso siguiente no ejecutaria
// https://www.youtube.com/watch?v=r5BWIXiQqH8
func main() {
    fmt.Println("Starting...")
    bandaTransp := make(chan string) // channel con capacidad 0

    go func() { // goroutine anonima para que ponga y evitar DeadLock
        fmt.Println("go func - enviando pieza 1 a banda..")
        bandaTransp <- "ThePieza1" // poniendo en channel
        fmt.Println("go func - pieza 1 enviada")
    }()

    fmt.Println("main - esperando pieza 1 en banda...")
    pieza1 := <- bandaTransp // sacando del channel
    fmt.Println("main - pieza 1 recibida: ", pieza1)
}
