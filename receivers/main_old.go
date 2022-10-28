package main

import "fmt"

func main() {
    SayHiTo("John")
}

func SayHiTo(name string) {
    fmt.Println("Hi:", name)
}
