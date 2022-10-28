package main

import "fmt"

type Player struct {
    Points int
}

func main() {
    player1 := Player{
        Points: 0,
    }

    fmt.Printf("Player address: %p\n", &player1)
    fmt.Println("Player points:", player1.Points)

    player1.AddPoints(100)

    // Points should be 100
    fmt.Println("Player Points:", player1.Points)

    player1.AddPoints(10)
    // Points should be 110
    fmt.Println("Player Points:", player1.Points)
}

//func (r Player) AddPoints(points int) { // doesn't work
func (r *Player) AddPoints(points int) { // pointer receiver
    //fmt.Printf("Player address: %p\n", &r)
    fmt.Printf("Player address: %p\n", r)
    r.Points += points
}
