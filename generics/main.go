package main

import (
    "fmt"
    "reflect"
    "golang.org/x/exp/constraints"
)

func main() {
    // Generics allow types to be parameters
    // Go 1.18+
    // main features:
    // 1. Type Parameter (with constraint)
    // 2. Type Inference
    // 3. Type Set

    fmt.Println(minInt(1, 2))
    fmt.Println(minFloat64(0.1, 0.2))

    // Generic
    x := GMin[int](2, 3)
    fmt.Println(x)

    //fmt.Println(min[int](1, 2))
    fmt.Println(min[float64](0.1, 0.2))
    fmt.Println(min(0.1, 0.2))

    type superFloat float64
    var sf superFloat = 0.3
    fmt.Println(min(sf, 0.2))

    fmt.Println(min[int](1, 2))
}

func GMin[T constraints.Ordered](x, y T) T {
    if x < y {
        return x
    }
    return y
}

type minTypes interface {
    ~float64 | int
}

// 1. Type parameter (with constraint)
// func min[T interface{}](a T, b T) T {
// func min[T any](a T, b T) T { // no constraint
// func min[T int](a T, b T) T {
// func min[T float64](a T, b T) T {
// func min[T ~float64](a T, b T) T { // 2. Type Inference
func min[T minTypes](a T, b T) T { // 3 Type Set
    fmt.Println("Generic: min a=",reflect.TypeOf(a))
    if a < b { // Don't work with "any"
        return a
    }
    return b
}

// Problem
func minInt(a int, b int) int {
    if a < b {
        return a
    }
    return b
}
func minFloat64(a float64, b float64) float64 {
    if a < b {
        return a
    }
    return b
}
