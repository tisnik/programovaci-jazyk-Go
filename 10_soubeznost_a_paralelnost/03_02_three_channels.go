// Gorutiny a kanály v jazyce Go
//
// - konstrukce select-case
// - čekání na data poslaná do libovolného kanálu

package main

import "fmt"

func f(c1 chan int, c2 chan int, c3 chan int) {
	// čekání na data poslaná do jednoho z kanálů
	select {
	case x := <-c1:
		fmt.Println("data from channel C1:", x)
	case y := <-c2:
		fmt.Println("data from channel C2:", y)
	case z := <-c3:
		fmt.Println("data from channel C3:", z)
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
