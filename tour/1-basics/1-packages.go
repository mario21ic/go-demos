package main

import (
    "fmt"
    "math/rand"
)

func main() {
    rand.Seed(1000) // to avoid a deterministic state with return 1
    fmt.Println("My favorite number is", rand.Intn(10))
}
