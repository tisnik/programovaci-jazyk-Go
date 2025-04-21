// Gorutiny a kanály v jazyce Go
//
// - konstrukce select-case
// - zápis do kanálů v konstrukci select-case

package main

import (
	"fmt"
	"time"
)

func worker(id int, channel chan int) {
	for value := range channel {
		fmt.Printf("Worker %d received value %d\n", id, value)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("Channel is closed\n")
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go worker(1, ch1)
	go worker(2, ch2)

	for i := 0; i < 10; i++ {
		select {
		case ch1 <- 1:
			fmt.Println("Sent value to channel 1")
		case ch2 <- 2:
			fmt.Println("Sent value to channel 2")
		}
	}
	close(ch1)
	close(ch2)
}
