// Základy programovacího jazyka Go
//
// - unární operátory

package main

import "fmt"

func main() {
	// unární operátory + a -
	var i int = 42
	fmt.Println(+i)
	fmt.Println(-i)

	// unární operátor ^
	i = 0
	fmt.Println(^i)
	i++
	fmt.Println(^i)

	// unární operátor !
	var b bool = false
	fmt.Println(!b)

	// unární operátor &
	fmt.Println(&i)

	// unární operátory & a *
	var ptrI *int = &i
	fmt.Println(*ptrI)

	// unární operátor <-
	var channel chan int
	channel = make(chan int)
	<-channel
}
