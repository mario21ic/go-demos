/*
A sender can close a channel to indicate that no more values will be sent.
Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

v, ok := <-ch
ok is false if there are no more values to receive and the channel is closed.

The loop for i := range c receives values from the channel repeatedly until it is closed.

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren't like files; you don't usually need to close them.
Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
*/

package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x // assing the x to the var channel c
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10) // channel with type int and up to 10 items as capacity
	// go fibonacci(cap(c), c) // run a goroutine with 10 as capacity and the channel
	go fibonacci(5, c) // run the goroutine with 10 as capacity and the channel
	x := 0
	for i := range c { // c has the value from func fibonacci
		// fmt.Println(i)
		fmt.Printf("%d - %d\n", x, i)
		x += 1
	}
}
