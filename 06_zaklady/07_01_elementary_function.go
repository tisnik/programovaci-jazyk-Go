// Základy programovacího jazyka Go
//
// - deklarace funkce bez parametrů a bez návratové hodnoty

package main

import "fmt"

// Deklarace funkce bez parametrů a bez návratové hodnoty
func sayHello() {
	fmt.Println("Hello, world!")
}

func main() {
	// zavolání funkce
	sayHello()
}
