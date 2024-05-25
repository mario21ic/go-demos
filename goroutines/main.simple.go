package main

import (
    "fmt"
    "time"
)

func Hello() {
    fmt.Println("Hola")
}

func main() {
    fmt.Println("## Starting")

    go Hello()
    time.Sleep(1 * time.Second)

    fmt.Println("## Finished")
}
