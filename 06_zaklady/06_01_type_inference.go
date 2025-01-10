// Základy programovacího jazyka Go
//
// - automatické odvození typu proměnné

package main

import "fmt"

func main() {
	counter := 1
	fmt.Printf("%T %v", counter, counter)
}
