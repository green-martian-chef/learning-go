// Channels
//
// Channels provide a way for two goroutines to communicate with each other and
// synchronize their execution. You can send values into channels from one
// goroutine and receive those values into another one.
package main

import "fmt"
import "time"

func sender(n chan int) {
	for i := 0; i < 10; i++ {

		// Send a value into a channel using `channel <- value`
		n <- i
	}
}

// When you use a channel to synchronize two goroutines, the goroutine which
// sends a value wait until the other goroutine is ready to receive the value.
// You are blocking the execution.
func receiver(n chan int) {
	for {
		// Receive a value from a channel with ` <- channel`
		num := <-n
		fmt.Println(num)
	}
}

func pinger(s chan string) {
	for i := 0; i < 10; i++ {
		s <- "ping"
	}
}

func ponger(s chan string) {
	for i := 0; i < 11; i++ {
		s <- "pong"
	}
}

func printer(s chan string) {
	for {
		str := <-s
		fmt.Println(str)
		time.Sleep(time.Second * 1)
	}
}

func outcome(o chan<- string) {
	o <- "Message"
}
func income(i <-chan string) {
	fmt.Println(<-i)
}

func main() {
	// Create a new channel with make(chan type)
	m := make(chan string)

	go pinger(m)

	// "ping" is passed to Println() through the channel m
	fmt.Println(<-m)
	fmt.Scanln()

	// Channels are unbuffered. They only accept sends (chan <-) if there is a
	// corresponding receive (<- chan) ready to receive the sent value. Buffered
	// channels accept a limited number of values without a corresponding receiver
	// for those values
	buffer := make(chan string, 2)

	buffer <- "first"
	buffer <- "second"

	fmt.Println(<-buffer)
	fmt.Println(<-buffer)

	s := make(chan string)

	go pinger(s)
	go ponger(s)
	go printer(s)

	fmt.Scanln()

	// Channel Synchronization
	n := make(chan int)

	go sender(n)
	go receiver(n)

	fmt.Scanln()
	fmt.Println("done")

	// Channel Directions
	// It is possible to specify if a channel is meant to only send or receive
	// values. By default, channels are bi-directional. This specificity increases
	// the type-safety of the program.

	o := make(chan string)

	go outcome(o)
	go income(o)
	fmt.Scanln()
}
