package main

import (
	"fmt"
)

type Speaker interface {
	Say(string)
}

type gopher struct {
	name string
}

// Implementacion de Say en gopher
func (g *gopher) Say(message string) {
	fmt.Printf("THe gopher %s says: %s", g.name, message)
}

func callSpeaker(s Speaker) {
	s.Say("Hello World!!\n")
}

func main() {
	g := &gopher{"Jack"}
	callSpeaker(g)

	b := &gopher{"Blackie"}
	callSpeaker(b)
}
