// Gorutiny a kanály v jazyce Go
//
// - konstrukce select-case
// - čekání na data poslaná do libovolného kanálu

package main

import "fmt"

func f(c1 chan int, c2 chan int, c3 chan int) {
	// čekání na data poslaná do jednoho z kanálů
	select {
	case <-c1:
		fmt.Println("C1")
	case <-c2:
		fmt.Println("C2")
	case <-c3:
		fmt.Println("C3")
	}
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	go f(c1, c2, c3)

	// data pošleme do druhého kanálu
	c2 <- 42
}
