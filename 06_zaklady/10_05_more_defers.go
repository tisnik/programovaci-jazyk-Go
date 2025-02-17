// Základy programovacího jazyka Go
//
// - větší množství konstrukcí defer
// - ověření pořadí volání

package main

import "fmt"

func onFinish(i int) {
	fmt.Printf("Defer #%2d\n", i)
}

func main() {
	// větší množství konstrukcí defer
	defer onFinish(1)
	defer onFinish(2)
	defer onFinish(3)
	defer onFinish(4)
	fmt.Println("Finishing main() function")
}
