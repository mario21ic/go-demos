package main

import (
	"fmt"
	"time"
)

func main() {
	// sayWithSleep(1)
	go sayWithSleep(1)
	say(2)

	time.Sleep(3 * time.Second)
}

func say(i int) {
	fmt.Printf("La concurrencia en GO es chevere: %d\n", i)
}

func sayWithSleep(i int) {
	time.Sleep(2 * time.Second)
	say(i)
}
