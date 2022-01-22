package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("go_errors.txt")
	if err != nil {
		log.Fatal(err)
	}
}
