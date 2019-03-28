// `select` lets you wait on multiple channel operations.
// `select` picks the first channel that is ready and receives from it (or sends
// to it). If more than one channel are ready then it picks randomly. If none
// are ready, it blocks until one becomes available.
package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	go func() {
		for {
			select {
			case m1 := <-c1:
				fmt.Println("received", m1)
			case m2 := <-c2:
				fmt.Println("received", m2)
			}
		}
	}()
	fmt.Scanln()
}
